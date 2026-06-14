<script setup>
import { ref, onMounted } from "vue";
import { Events } from "@wailsio/runtime";
import JogDimple from "./JogDimple.vue";

const wheelPos = ref(0);
const FULL_ROTATION = 1474560; // 360 degrees * 4096 units/degree

onMounted(() => {
    // Listen for the Wails3 event
    Events.On("jogWheelMoved", (event) => {
        // 1. Accumulate the delta (distance moved) to the current position
        let newPos = wheelPos.value + event.data;

        // 2. Wrap the position around the 360-degree mark.
        // This prevents the number from growing infinitely large over time.
        // We add FULL_ROTATION before the final modulo to correctly handle negative deltas
        // (since JS modulo of a negative number returns a negative number).
        wheelPos.value =
            ((newPos % FULL_ROTATION) + FULL_ROTATION) % FULL_ROTATION;
    });
});
</script>

<template>
    <div class="jog-wheel">
        <!-- Pass the reactive wheelPos down to the child -->
        <JogDimple :wheel-pos="wheelPos" />
    </div>
</template>

<style scoped>
.jog-wheel {
    background: radial-gradient(circle, #3d4e60 0%, #2d3e50 70%, #1b2636 100%);
    border: 3px solid #4d5e70;
    border-radius: 50%;

    /* Grid placement (adjust column/row to match your layout) */
    grid-column: 17 / span 6;
    grid-row: 7 / span 6;

    aspect-ratio: 1 / 1;
    align-self: center;
    justify-self: center;
    width: 100%;
    max-width: 150px;
    max-height: 150px; /* Changed from 300px to 150px to respect the 1:1 aspect ratio */
    box-shadow:
        inset 0 0 20px rgba(0, 0, 0, 0.5),
        0 0 10px rgba(74, 158, 255, 0.2);

    /* REQUIRED: Allows the child component to position itself absolutely */
    position: relative;
    overflow: hidden; /* Keeps the dimple cleanly clipped inside the circular border */
}
</style>
