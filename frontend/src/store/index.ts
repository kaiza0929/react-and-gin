import { createStore } from "redux";

type TestLog = {

}

var store = createStore((state = [], action: {type: string, payload: TestLog[]}) => {
    switch (action.type) {
        default:
            return state;
    }
});

export default store;