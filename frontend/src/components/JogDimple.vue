<script setup>
import { computed } from "vue";

const props = defineProps({
    wheelPos: {
        type: Number,
        required: true,
    },
});

// 4096 units = 1 degree
const angle = computed(() => {
    // Calculate degrees.
    // We use modulo 360 so the number doesn't grow infinitely large
    // and to prevent floating-point precision issues over long sessions.
    return (props.wheelPos / 4096) % 360;
});
</script>

<template>
    <!-- This wrapper covers the whole dial and rotates based on the wheel position -->
    <div class="dimple-wrapper" :style="{ transform: `rotate(${angle}deg)` }">
        <!-- The dimple is pinned to the top edge of the wrapper -->
        <div class="dimple"></div>
    </div>
</template>

<style scoped>
.dimple-wrapper {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    pointer-events: none; /* Allows clicks to pass through to the dial if needed */
}

.dimple {
    position: absolute;
    /* Pin to the top edge. Adjust 'top' to move it further in/out.
       10px places it nicely on the inner edge of a 150px dial. */
    top: 10px;
    left: 50%;
    width: 14px;
    height: 14px;

    /* Creates a realistic 3D "dimpled" or indented look */
    background: radial-gradient(circle at 30% 30%, #ffffff, #8899aa);
    border-radius: 50%;
    transform: translateX(-50%); /* Centers the dimple horizontally */
    box-shadow:
        inset 0 2px 4px rgba(0, 0, 0, 0.6),
        0 1px 2px rgba(255, 255, 255, 0.2);
}
</style>
