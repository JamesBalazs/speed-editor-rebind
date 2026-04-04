<script setup>
import { ref, onMounted, computed } from "vue";
import { Events } from "@wailsio/runtime";
import { store } from "../store.js";
// import { GreetService } from "../../bindings/changeme";

const props = defineProps({
    id: Number,
    col: Number,
    colSpan: Number,
    row: Number,
    text: String,
    subText: String,
    led: Number,
    jogLed: Number,
});

const isActive = ref(false);
let activeTimeout = null;

onMounted(() => {
    Events.On(`keyPress-${props.id}`, (event) => {
        // Highlight the pressed keys
        pressed(event.data);

        // Log for debugging
        console.log(`Key pressed: ${props.id}, `, event.data);
    });
});

function pressed(data) {
    // Clear any pending timeout if key is pressed again
    if (activeTimeout) {
        clearTimeout(activeTimeout);
    }
    isActive.value = true;
    activeTimeout = setTimeout(() => {
        isActive.value = false;
    }, 250);
}

const formattedText = computed(() => {
    return props.text.replace(/\\n/g, "<br>");
});

const formattedSubText = computed(() => {
    return props.subText.replace(/\\n/g, "<br>");
});

const isSelected = computed(() => {
    return store.selectedKey == props.id;
});
</script>

<template>
    <div
        :class="{ key: true, active: isActive, selected: isSelected }"
        :id="`key-${id}`"
        :data-id="id"
        :style="`grid-column: ${col} / span ${colSpan}; grid-row: ${row} / span 2`"
        @click="store.selectKey(id)"
    >
        <div class="led-container" v-if="led != 0 || jogLed != 0">
            <span class="led"></span>
        </div>

        <span class="key-text" v-if="text !== ''" v-html="formattedText"></span>
        <span
            class="key-subtext"
            v-if="subText !== ''"
            v-html="formattedSubText"
        ></span>
    </div>
</template>

<style scoped>
.key {
    background: linear-gradient(145deg, #2d3e50, #1b2636);
    border: 1px solid #3d4e60;
    border-radius: 6px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 4px;
    cursor: pointer;
    transition: all 0.1s ease;
    position: relative;
    user-select: none;
    /* Fill the grid cell completely */
    width: 100%;
    height: 100%;
    box-sizing: border-box;
}

.key:hover {
    background: linear-gradient(145deg, #3d4e60, #2d3e50);
    border-color: #4d5e70;
}

.key.active {
    background: linear-gradient(145deg, #4a9eff, #2d6e9f);
    border-color: #6eb4ff;
    box-shadow: 0 0 15px rgba(74, 158, 255, 0.5);
    transform: translateY(-2px);
}

.key-text {
    font-size: 10px;
    font-weight: 600;
    text-align: center;
    color: #e0e0e0;
    line-height: 1.1;
    /* Handle line breaks */
    white-space: normal;
    word-break: break-word;
}

.key-text br {
    line-height: 1.1;
}

.key-subtext {
    font-size: 8px;
    color: #8899aa;
    text-align: center;
    margin-top: 2px;
    /* Handle line breaks */
    white-space: normal;
    word-break: break-word;
    line-height: 1.1;
}

.key-name {
    font-size: 7px;
    color: #556677;
    position: absolute;
    bottom: 3px;
    left: 4px;
}

.led {
    display: inline-block;
    border: 1px solid #333;
    padding: 2px 7px;
    background-color: #001;
}

.led-container {
    margin-top: -8px;
}

.led-lit {
    background-color: rgba(255, 38, 54, 1);
}

.key.selected {
    box-shadow: 0 0 0 3px #fbbf24;
}
</style>
