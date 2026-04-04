import { reactive } from "vue";

export const store = reactive({
    selectedKey: null,
    selectedLed: null,
    selectedJogLed: null,
    selectKey(id, ledId, jogLedId) {
        this.selectedKey = id;
        this.selectedLed = ledId;
        this.selectedJogLed = jogLedId;
    },
});
