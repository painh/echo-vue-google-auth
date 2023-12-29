import { InjectionKey } from "vue";
import { createStore, useStore as baseUseStore, Store } from "vuex";

export interface State {
    count: number;
}

export const key: InjectionKey<Store<State>> = Symbol();
// vue 컴포넌트에서 store 내부의 state에 대한 타입을 추론하게 해준다.

export default createStore<State>({
    state() {
        return {
            count: 0,
        };
    },
    mutations: {
        increment(state : State) {
            state.count++;
        },
    },
});

export const useStore = () => {
    return baseUseStore(key);
};