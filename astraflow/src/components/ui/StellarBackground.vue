<template>
  <div class="stellar-background">
    <!-- Nebula Background -->
    <div class="nebula-bg"></div>

    <!-- Stellar Field -->
    <div class="stellar-field">
      <div
        v-for="star in stars"
        :key="star.id"
        class="star"
        :style="star.style"
      ></div>
    </div>

    <!-- Aurora Borealis Effect -->
    <div class="aurora-layer"></div>

    <!-- Cosmic Particles -->
    <div class="particle-field">
      <div
        v-for="particle in particles"
        :key="particle.id"
        class="particle"
        :style="particle.style"
      ></div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const stars = ref([])
const particles = ref([])

onMounted(() => {
  // Generate random stars
  const starArray = []
  for (let i = 0; i < 100; i++) {
    starArray.push({
      id: i,
      style: {
        left: `${Math.random() * 100}%`,
        top: `${Math.random() * 100}%`,
        width: `${Math.random() * 3}px`,
        height: `${Math.random() * 3}px`,
        animationDelay: `${Math.random() * 10}s`,
        animationDuration: `${10 + Math.random() * 20}s`
      }
    })
  }
  stars.value = starArray

  // Generate cosmic particles
  const particleArray = []
  for (let i = 0; i < 20; i++) {
    particleArray.push({
      id: i,
      style: {
        left: `${Math.random() * 100}%`,
        animationDelay: `${Math.random() * 15}s`,
        animationDuration: `${15 + Math.random() * 25}s`
      }
    })
  }
  particles.value = particleArray
})
</script>

<style scoped>
.stellar-background {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
  overflow: hidden;
}

/* Nebula Background */
.nebula-bg {
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(
    ellipse at center,
    rgba(59, 130, 246, 0.1) 0%,
    rgba(139, 92, 246, 0.05) 30%,
    rgba(236, 72, 153, 0.03) 60%,
    transparent 100%
  );
  animation: nebulaFlow 20s ease-in-out infinite;
  opacity: var(--nebula-opacity);
}

/* Light Theme Nebula */
[data-theme="light"] .nebula-bg {
  background: radial-gradient(
    ellipse at center,
    rgba(59, 130, 246, 0.06) 0%,
    rgba(139, 92, 246, 0.03) 30%,
    rgba(236, 72, 153, 0.02) 60%,
    transparent 100%
  );
}

/* Stellar Field */
.stellar-field {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.star {
  position: absolute;
  background: white;
  border-radius: 50%;
  animation: stellarField linear infinite;
  box-shadow: 0 0 6px rgba(255, 255, 255, 0.8);
  opacity: var(--particle-opacity);
}

.star:nth-child(3n) {
  background: var(--color-cosmic-cyan);
  box-shadow: 0 0 10px var(--color-cosmic-cyan);
}

.star:nth-child(5n) {
  background: var(--color-nebula-purple);
  box-shadow: 0 0 10px var(--color-nebula-purple);
}

/* Light Theme Stars */
[data-theme="light"] .star {
  background: rgba(59, 130, 246, 0.6);
  box-shadow: 0 0 4px rgba(59, 130, 246, 0.4);
}

[data-theme="light"] .star:nth-child(3n) {
  background: rgba(6, 182, 212, 0.6);
  box-shadow: 0 0 6px rgba(6, 182, 212, 0.4);
}

[data-theme="light"] .star:nth-child(5n) {
  background: rgba(139, 92, 246, 0.6);
  box-shadow: 0 0 6px rgba(139, 92, 246, 0.4);
}

/* Aurora Layer */
.aurora-layer {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    45deg,
    transparent 30%,
    rgba(59, 130, 246, 0.05) 50%,
    rgba(139, 92, 246, 0.05) 70%,
    transparent 90%
  );
  animation: auroraWave 30s linear infinite;
  opacity: var(--nebula-opacity);
}

/* Light Theme Aurora */
[data-theme="light"] .aurora-layer {
  background: linear-gradient(
    45deg,
    transparent 30%,
    rgba(59, 130, 246, 0.03) 50%,
    rgba(139, 92, 246, 0.03) 70%,
    transparent 90%
  );
}

/* Particle Field */
.particle-field {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.particle {
  position: absolute;
  width: 2px;
  height: 2px;
  background: var(--color-stellar-blue);
  border-radius: 50%;
  animation: particleFloat linear infinite;
  box-shadow: 0 0 10px var(--color-nebula-glow);
  opacity: var(--particle-opacity);
}

.particle:nth-child(2n) {
  background: var(--color-nebula-purple);
  box-shadow: 0 0 10px rgba(139, 92, 246, 0.8);
}

.particle:nth-child(3n) {
  background: var(--color-aurora-pink);
  box-shadow: 0 0 10px var(--color-aurora-glow);
}

/* Light Theme Particles */
[data-theme="light"] .particle {
  background: rgba(59, 130, 246, 0.4);
  box-shadow: 0 0 6px rgba(59, 130, 246, 0.3);
}

[data-theme="light"] .particle:nth-child(2n) {
  background: rgba(139, 92, 246, 0.4);
  box-shadow: 0 0 6px rgba(139, 92, 246, 0.3);
}

[data-theme="light"] .particle:nth-child(3n) {
  background: rgba(236, 72, 153, 0.4);
  box-shadow: 0 0 6px rgba(236, 72, 153, 0.3);
}
</style>