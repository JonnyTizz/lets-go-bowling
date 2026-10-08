---
theme: default
title: Let's Go Bowling
transition: slide-left
layout: center
clicks: 1
---

<BowlingIntro :play="$clicks > 0" />

<!--
TL;DR: Look at Go, then write some.

- Title on its own first.
- Very short intro. We're going to look at Go,
  then we're going to actually write some.

[click] Gophers fly in, ball follows after a beat. Strike. 🎳
-->

---

# Why Go?

<div class="why-go">

<v-clicks>

- 🚀 **Pretty dang speedy**
- 🏗️ **Compiles fast**
- 📦 **Deploys simply**
- 🧠 **Simple language**
- 🛠️ **Opinionated tooling**

</v-clicks>

</div>

<style>
.why-go {
  font-size: 1.6rem;
  line-height: 2.6rem;
  margin-top: 1.5rem;
}
</style>

<!--
TL;DR: It's the combo, not any one feature. Boring on purpose.

- Don't explain Go feature-by-feature.
- Talk around the combination: it's not just runtime speed,
  the whole development/deployment experience is deliberately
  boring/simple.
-->

---

# A useful middle ground

<div class="spectrum">
  <div class="track" />
  <div class="end left">JS / TS</div>
  <div class="end right">Rust / C++</div>
  <div v-click="1" class="go-drop">Go</div>
</div>

<p v-click="2" class="text-center available">🤷‍♂️ Other languages are available</p>

<style>
.spectrum {
  position: relative;
  height: 11rem;
  margin: 3rem 2rem 1rem;
}
.track {
  position: absolute;
  left: 0;
  right: 0;
  top: 7rem;
  height: 4px;
  border-radius: 2px;
  background: linear-gradient(90deg, #f7df1e, #00add8 45%, #00add8 55%, #D34516);
}
.end {
  position: absolute;
  top: 8rem;
  font-size: 1.4rem;
  font-weight: 600;
}
.end.left { left: 0; }
.end.right { right: 0; }
.go-drop {
  position: absolute;
  left: 50%;
  top: 4.4rem;
  translate: -50% 50%;
  padding: 0.2rem 1.2rem;
  border-radius: 999px;
  background: #00add8;
  color: white;
  font-size: 1.6rem;
  font-weight: 700;
}
.go-drop.slidev-vclick-target {
  transition: transform 700ms cubic-bezier(0.34, 1.56, 0.64, 1), opacity 200ms ease;
}
.go-drop.slidev-vclick-hidden {
  transform: translateY(-6rem);
}
.available {
  margin-top: 6rem;
  font-size: 1.6rem;
}
</style>

<!--
TL;DR: Compiled-language perks without the Rust/C++ complexity. Simplicity is the point.

- Hand-wavey on purpose, not a benchmark.

[click] Go drops in.

- Go gives us a bunch of "proper compiled language" properties
  without taking us all the way into the complexity of Rust/C++.
- Tradeoff: Go intentionally leaves lots of features out.
  That simplicity is part of the design.

[click] Other languages are available.
-->

---

# Enough slides.

## Let's write some Go.

<!--
TL;DR: You drive, I type. JS brain → Go code, compiler and tests keep us honest.

Switch to editor here. Audience rules:
- They tell me what to do.
- Translate it into Go.
- Let compiler/tests guide us.
- Stop and investigate interesting Go-isms as they appear.
-->

---
layout: center
---

<BowlingTitle />

<!--
TL;DR: Holding slide.

Use it to switch windows, reset the exercise, take questions, etc.
-->
