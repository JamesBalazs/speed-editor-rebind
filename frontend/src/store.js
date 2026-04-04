import { reactive } from "vue";

export const store = reactive({
    selectedKey: null,
    selectKey(id) {
        this.selectedKey = id;
    },
});
