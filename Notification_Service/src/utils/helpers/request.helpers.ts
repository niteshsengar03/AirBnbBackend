import { AsyncLocalStorage } from "async_hooks";

type AsyncLocalStorageType =   {
    correlationId:string;
}
// create a instance of async storage
export const asyncLocalStorage = new AsyncLocalStorage<AsyncLocalStorageType>();


export const getCorrelationId = () =>{
    const asyncStore = asyncLocalStorage.getStore();
    return asyncStore?.correlationId || "unkown-error-while-creating-correlation-id";
}

