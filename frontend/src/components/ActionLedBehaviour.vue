<script setup>
import { ref, onMounted, computed, watch } from "vue";
import { store } from "../store.js";
import { SpeedEditorService } from "../../bindings/github.com/JamesBalazs/speed-editor-rebind";

const splitter = ref(20);
const tab = ref("none");

const selectedKey = computed(() => {
    return store.selectedKey;
});

const selectedLed = computed(() => {
    return store.selectedLed;
});

const selectedJogLed = computed(() => {
    return store.selectedJogLed;
});

watch(tab, (newValue, oldValue) => {
    if (!store.selectedHasAnyLed()) {
        return;
    }

    SpeedEditorService.SetKeyLedBehaviour(store.selectedKey, newValue);
});
</script>

<template>
    <q-splitter>
        <template v-slot:before v-model="splitter">
            <q-tabs v-model="tab" vertical class="text-teal">
                <q-tab name="none" label="None" />
                <q-tab name="flash" label="Flash" />
                <q-tab name="latch" label="Latch" />
            </q-tabs>
        </template>

        <template v-slot:after>
            <q-tab-panels
                v-model="tab"
                animated
                swipeable
                vertical
                dark
                transition-prev="jump-up"
                transition-next="jump-up"
            >
                <q-tab-panel name="none">
                    <p>No LED on keypress.</p>
                </q-tab-panel>

                <q-tab-panel name="flash">
                    <p>LED flashes on keypress.</p>
                </q-tab-panel>

                <q-tab-panel name="latch">
                    <p>LED toggles on keypress.</p>
                </q-tab-panel>
            </q-tab-panels>
        </template>
    </q-splitter>
</template>

<style scoped></style>
