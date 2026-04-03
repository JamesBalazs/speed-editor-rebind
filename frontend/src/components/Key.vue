<script setup>
import { ref, onMounted, computed } from "vue";
import { Events } from "@wailsio/runtime";
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

onMounted(() => {
    Events.On(`keyPress-${props.id}`, (event) => {
        // Highlight the pressed keys
        pressed(event.data);

        // Log for debugging
        console.log(`Key pressed: ${props.id}, `, event.data);
    });
});

function pressed(data) {
    console.log("pressed called with: ", data);
}

const formattedText = computed(() => {
    return props.text.replace(/\\n/g, "<br>");
});

const formattedSubText = computed(() => {
    return props.subText.replace(/\\n/g, "<br>");
});
</script>

<template>
    <div
        class="key"
        :id="`key-${id}`"
        :data-id="id"
        :style="`grid-column: ${col} / span ${colSpan}; grid-row: ${row} / span 2`"
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
