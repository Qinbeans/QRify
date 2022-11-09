import { writable, type Writable } from 'svelte/store';
export const active: Writable<number> = writable(0);