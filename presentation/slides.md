---
theme: default
title: Let's Go Bowling
transition: slide-left
layout: center
---

<BowlingTitle />

<!--
TL;DR: this is a real brief look at Go, meant to whet your appetite
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
Pretty dang speedy & compiles fast
- built to be a replacement at Google for C++
- actually conceived by 3 engineers while they were waiting for their code to compile

Deploys simply
- native binaries can be built for any platform from any platform

Simple language
- Intentionally left features out when it was being built
- Generally one way to do things
- Really complete std lib - e.g. the included HTTP server is featured enough for prod!

Opinionated tooling
- compiler breaks if vars are unused
- formatter built in
- testing and framework included
- version and dependencies built in
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
Compiled-language perks without the Rust/C++ complexity. 

[click] Go drops in.

- Go gives us a bunch of "proper compiled language" properties, like the speed 
  and the compile time checks, without taking us all the way into the complexity
  of Rust/C++ with manual memory management.
- The goal was to make a language where anyone could look at any project and
  understand it without having to get to grips with each teams specific styles
- Tradeoff: Go intentionally leaves lots of features out, and can feel repetitive

[click] Other languages are available.
-->

---

# Enough slides.

## Let's write some Go.

<!--
TL;DR: You drive, I type. JS brain → Go code, compiler and tests keep us honest.

Audience rules:
- They tell me what to do.
- Translate it into Go.
- Let compiler/tests guide us.
- Stop and investigate interesting Go-isms as they appear.
-->

---
layout: center
clicks: 1
---

<BowlingIntro :play="$clicks > 0" />

<!--
TL;DR: The bowling joke, then over to the editor.

- Title on its own first.

[click] Gophers lowered in, ball follows after a beat. Strike. 🎳

Switch to editor here. Doubles as a holding slide too.
-->
