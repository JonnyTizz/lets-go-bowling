<script setup lang="ts">
import { useIsSlideActive } from '@slidev/client'
import { computed, onBeforeUnmount, ref, watch } from 'vue'
import gopher from '../assets/gopher.svg'

// ---------------------------------------------------------------------------
// World: inches. x runs down the lane towards the pit, y is to the bowler's
// left, z is up. The head pin sits at the origin.
// ---------------------------------------------------------------------------
const PIN_SPACING = 12
const ROW_DEPTH = PIN_SPACING * Math.sqrt(3) / 2
const PIN_H = 15
const PIN_R = 2.4
const BALL_R = 4.3

// Camera: crouched ~2ft off the floor near the pin end, having walked round
// to the right of the rack.
const CAMERA = {
  orbit: 52, // degrees round to the right from straight down the lane
  distance: 78,
  height: 24,
  lookHeight: 7,
}

// Scene box (px) the rack is framed inside
const SCENE_W = 640
const SCENE_H = 300
const RACK_FILL = 0.44 // fraction of the scene width the standing rack takes up (smaller = further back)
const RACK_CENTRE = { x: 0.66, y: 0.45 } // where the standing rack sits in the scene, as fractions

// Timings (ms, from when `play` turns on)
const LOWER_MS = 1400 // the whole rack being lowered onto the lane
const LOWER_FROM = 480 // px above its spot that the rack starts (above the top of the slide)
const BALL_PAUSE = 600 // after the last gopher lands
const ROLL_MS = 220 // from off the slide to impact, at a constant speed
const BALL_KICKBACK = 4 // in, how far the ball bounces back
const BALL_HOP = 3 // in, height of the first bounce
const PIN_FALL_MS = 260
const SLIDE_DECAY = 4 // how quickly a knocked gopher slows down

type V3 = [number, number, number]
const sub = (a: V3, b: V3): V3 => [a[0] - b[0], a[1] - b[1], a[2] - b[2]]
const dot = (a: V3, b: V3) => a[0] * b[0] + a[1] * b[1] + a[2] * b[2]
const cross = (a: V3, b: V3): V3 => [a[1] * b[2] - a[2] * b[1], a[2] * b[0] - a[0] * b[2], a[0] * b[1] - a[1] * b[0]]
const norm = (a: V3): V3 => {
  const l = Math.hypot(...a)
  return [a[0] / l, a[1] / l, a[2] / l]
}
const rand = (i: number) => {
  const s = Math.sin(i * 12.9898) * 43758.5453
  return s - Math.floor(s)
}

const pinsAtRest = [0, 1, 2, 3].flatMap(row =>
  Array.from({ length: row + 1 }, (_, k): V3 => [row * ROW_DEPTH, (k - row / 2) * PIN_SPACING, 0]),
)

// Camera basis
const target: V3 = [2 * ROW_DEPTH, 0, CAMERA.lookHeight]
const theta = CAMERA.orbit * Math.PI / 180
const eye: V3 = [target[0] - CAMERA.distance * Math.cos(theta), -CAMERA.distance * Math.sin(theta), CAMERA.height]
const forward = norm(sub(target, eye))
const right = norm(cross(forward, [0, 0, 1]))
const up = cross(right, forward)

function rawProject(p: V3) {
  const rel = sub(p, eye)
  const depth = dot(rel, forward)
  return { x: dot(rel, right) / depth, y: -dot(rel, up) / depth, depth }
}

// Frame the standing rack in the scene
const bounds = pinsAtRest.flatMap(p => [rawProject(p), rawProject([p[0], p[1], PIN_H])])
const minX = Math.min(...bounds.map(b => b.x))
const maxX = Math.max(...bounds.map(b => b.x))
const minY = Math.min(...bounds.map(b => b.y))
const maxY = Math.max(...bounds.map(b => b.y))
const focal = SCENE_W * RACK_FILL / (maxX - minX)
const offset = {
  x: SCENE_W * RACK_CENTRE.x - focal * (minX + maxX) / 2,
  y: SCENE_H * RACK_CENTRE.y - focal * (minY + maxY) / 2,
}

function project(p: V3) {
  const r = rawProject(p)
  return { x: offset.x + focal * r.x, y: offset.y + focal * r.y, scale: focal / r.depth, depth: r.depth }
}

// Ball: rolls straight at the head pin from wherever it's off the left of the slide
const contactX = -(BALL_R + PIN_R)
// Slide edges in scene coordinates (the slide is 980 x 552, scene sits roughly centred)
const SLIDE_LEFT = SCENE_W / 2 - 490
const SLIDE_RIGHT = SCENE_W / 2 + 490
const SLIDE_BOTTOM = SCENE_H + 70
let startX = contactX
while (startX > contactX - 400) {
  const p = project([startX, 0, BALL_R])
  if (p.depth < 10 || p.x + p.scale * BALL_R < SLIDE_LEFT - 20)
    break
  startX -= 1
}
const BALL_DELAY = LOWER_MS + BALL_PAUSE
const impactAt = BALL_DELAY + ROLL_MS
const END_MS = impactAt + 2500

// Each pin gets knocked away from the ball, with a little chaos
const pinMotion = pinsAtRest.map((p, i) => {
  // Spread sideways more than straight on, so the rack splays rather than all going in the pit
  const away = norm([p[0] - contactX + 4, 1.8 * p[1] + (rand(i + 1) - 0.5) * 10, 0])
  const spin = (rand(i + 2) - 0.5) * 0.8
  const dir: V3 = [away[0] * Math.cos(spin) - away[1] * Math.sin(spin), away[0] * Math.sin(spin) + away[1] * Math.cos(spin), 0]
  // How far it slides before stopping, cut short so it never leaves the slide
  let travel = 10 + rand(i + 3) * 26
  while (travel > 0) {
    const end = project([p[0] + dir[0] * travel, p[1] + dir[1] * travel, 0])
    const lying = end.scale * PIN_H // toppled, it sticks out about its own height
    if (end.x + lying < SLIDE_RIGHT - 20 && end.x - lying > SLIDE_LEFT + 20 && end.y < SLIDE_BOTTOM - 20)
      break
    travel -= 1
  }
  return {
    dir,
    travel,
    hitAt: impactAt + Math.hypot(p[0], p[1]) / 300 * 1000 + rand(i + 4) * 40,
  }
})

// ---------------------------------------------------------------------------
// Animation clock. Starts when `play` turns on (a click on the slide), and
// resets when it turns off or we leave the slide.
// ---------------------------------------------------------------------------
const props = defineProps<{ play: boolean }>()
const active = useIsSlideActive()
const running = computed(() => active.value && props.play)
const now = ref(0)
let raf = 0
let startedAt = 0

function tick(ts: number) {
  if (!startedAt)
    startedAt = ts
  now.value = ts - startedAt
  if (now.value < END_MS)
    raf = requestAnimationFrame(tick)
}

watch(running, (isRunning) => {
  cancelAnimationFrame(raf)
  now.value = 0
  startedAt = 0
  if (isRunning)
    raf = requestAnimationFrame(tick)
}, { immediate: true })
onBeforeUnmount(() => cancelAnimationFrame(raf))

// ---------------------------------------------------------------------------
// Per-frame state
// ---------------------------------------------------------------------------
const easeOut = (t: number) => 1 - (1 - t) ** 3
const easeInOut = (t: number) => t < 0.5 ? 4 * t ** 3 : 1 - (-2 * t + 2) ** 3 / 2

// How far the rack has been lowered, 0 → 1
const lowered = computed(() => Math.min(1, now.value / LOWER_MS))

const pins = computed(() => running.value ? pinsAtRest.map((rest, i) => {
  const m = pinMotion[i]
  const t = Math.max(0, now.value - m.hitAt) / 1000
  const slid = m.travel * (1 - Math.exp(-SLIDE_DECAY * t))
  const pos: V3 = [rest[0] + m.dir[0] * slid, rest[1] + m.dir[1] * slid, 0]

  const foot = project(pos)
  const fall = Math.min(1, (t * 1000 / PIN_FALL_MS) ** 2)
  const ahead = project([pos[0] + m.dir[0], pos[1] + m.dir[1], pos[2]])
  const sideways = Math.max(-1, Math.min(1, (ahead.x - foot.x) / (foot.scale || 1)))
  return {
    x: foot.x,
    y: foot.y,
    lift: -(1 - easeInOut(lowered.value)) * LOWER_FROM,
    z: Math.round(10000 - foot.depth * 10),
    height: foot.scale * PIN_H,
    angle: fall * 90 * sideways,
    squash: 1 - fall * (1 - Math.abs(sideways)) * 0.55,
    shadow: lowered.value ** 4,
  }
}) : [])

const ball = computed(() => {
  const t = now.value
  if (t < BALL_DELAY)
    return null
  let x: number
  let hop = 0
  if (t < impactAt) {
    x = startX + (contactX - startX) * (t - BALL_DELAY) / ROLL_MS
  }
  else {
    const s = (t - impactAt) / 1000
    x = contactX - BALL_KICKBACK * easeOut(Math.min(1, s / 0.45))
    hop = BALL_HOP * Math.abs(Math.sin(s * Math.PI / 0.22)) * Math.exp(-s / 0.12)
  }
  const centre = project([x, 0, BALL_R + hop])
  const floor = project([x, 0, 0])
  return {
    x: centre.x,
    y: centre.y,
    floorY: floor.y,
    z: Math.round(10000 - centre.depth * 10) + 1,
    size: centre.scale * BALL_R * 2,
    spin: (x - startX) / BALL_R * 180 / Math.PI,
    shadow: 1 - Math.min(0.6, hop / BALL_HOP),
  }
})
</script>

<template>
  <div class="bowling-intro">
    <BowlingTitle class="title" />

    <div class="scene" :style="{ width: `${SCENE_W}px`, height: `${SCENE_H}px` }">
      <template v-for="(pin, i) in pins" :key="i">
        <div
          class="shadow"
          :style="{
            left: `${pin.x}px`,
            top: `${pin.y}px`,
            width: `${pin.height * 0.55}px`,
            height: `${pin.height * 0.12}px`,
            opacity: 0.18 * pin.shadow,
          }"
        />
        <img
          :src="gopher"
          alt=""
          class="pin"
          :style="{
            left: `${pin.x}px`,
            top: `${pin.y + pin.lift}px`,
            height: `${pin.height}px`,
            zIndex: pin.z,
            transform: `translate(-50%, -100%) rotate(${pin.angle}deg) scale(-1, ${pin.squash})`,
          }"
        >
      </template>

      <template v-if="ball">
        <div
          class="shadow"
          :style="{
            left: `${ball.x}px`,
            top: `${ball.floorY}px`,
            width: `${ball.size * 0.9}px`,
            height: `${ball.size * 0.22}px`,
            opacity: 0.3 * ball.shadow,
          }"
        />
        <div
          class="ball"
          :style="{
            left: `${ball.x}px`,
            top: `${ball.y}px`,
            width: `${ball.size}px`,
            height: `${ball.size}px`,
            zIndex: ball.z,
            transform: `translate(-50%, -50%) rotate(${ball.spin}deg)`,
          }"
        >
          <span /><span /><span />
        </div>
      </template>
    </div>
  </div>
</template>

<style scoped>
.title {
  position: relative;
  z-index: 1;
}

.scene {
  position: relative;
  /* own stacking context, so the title can sit in front of the whole scene */
  z-index: 0;
  margin: 1rem auto 0;
}

.pin {
  position: absolute;
  max-width: none;
  transform-origin: 50% 100%;
}

.shadow {
  position: absolute;
  border-radius: 50%;
  background: #000;
  filter: blur(3px);
  transform: translate(-50%, -50%);
}

.ball {
  position: absolute;
  border-radius: 50%;
  background: radial-gradient(circle at 30% 30%, #5b5b6b, #15151c 70%);
}
.ball span {
  position: absolute;
  width: 13%;
  height: 13%;
  border-radius: 50%;
  background: #050507;
}
.ball span:nth-child(1) { top: 23%; left: 42%; }
.ball span:nth-child(2) { top: 37%; left: 62%; }
.ball span:nth-child(3) { top: 54%; left: 54%; width: 17%; height: 17%; }
</style>
