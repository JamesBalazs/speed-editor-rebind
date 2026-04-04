<script setup>
import { ref, onMounted, computed } from "vue";
import { store } from "../store.js";
// import { GreetService } from "../../bindings/changeme";

const selectedKey = computed(() => {
    return store.selectedKey;
});
</script>

<template>
    <q-card dark bordered flat class="toolbox">
        <q-toolbar class="glossy row">
            <q-toolbar-title v-if="selectedKey === null"
                >No Key Selected</q-toolbar-title
            >
            <q-toolbar-title v-else>Choose Action</q-toolbar-title>
        </q-toolbar>

        <q-tabs align="left" :class="{ disabled: selectedKey === null }">
            <q-tab name="button" label="Button"></q-tab>
            <q-tab name="led" label="LED Behaviour"></q-tab>
        </q-tabs>

        <q-separator dark></q-separator>
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
