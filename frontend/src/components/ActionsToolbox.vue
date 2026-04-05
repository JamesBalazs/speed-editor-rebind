<script setup>
import { ref, onMounted, computed, watch } from "vue";
import { store } from "../store.js";
// import { GreetService } from "../../bindings/changeme";

import ActionLedBehaviour from "./ActionLedBehaviour.vue";

const selectedKey = computed(() => {
    return store.selectedKey;
});

const hasLed = computed(() => {
    return store.selectedHasAnyLed();
});

watch(hasLed, (newValue, oldValue) => {
    if (newValue === false) {
        tab.value = "button";
    }
});

const tab = ref("button");
</script>

<template>
    <q-card dark bordered flat class="toolbox">
        <q-toolbar class="glossy row">
            <q-toolbar-title v-if="selectedKey === null"
                >No Key Selected</q-toolbar-title
            >
            <q-toolbar-title v-else>Choose Action</q-toolbar-title>
        </q-toolbar>

        <q-tabs
            v-model="tab"
            align="left"
            :class="{ disabled: selectedKey === null }"
        >
            <q-tab name="button" label="Button"></q-tab>
            <q-tab name="led" label="LED Behaviour" v-if="hasLed"></q-tab>
        </q-tabs>

        <q-separator dark></q-separator>

        <q-tab-panels dark v-model="tab" animated>
            <q-tab-panel name="button"> btn </q-tab-panel>
            <q-tab-panel name="led" v-if="hasLed">
                <ActionLedBehaviour></ActionLedBehaviour>
            </q-tab-panel>
        </q-tab-panels>
    </q-card>
</template>

<style scoped>
.toolbox {
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    width: 100vw;
    height: 300px;
    border-radius: 0;
    z-index: 1000;
}
</style>
