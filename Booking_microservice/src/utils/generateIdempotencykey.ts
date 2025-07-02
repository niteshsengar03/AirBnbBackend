import {v4 as uuidv4} from 'uuid'

export function generateIdempotencykey():string{
    return uuidv4();
}