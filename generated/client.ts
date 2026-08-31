// Code generated from docs/openapi.json by `make generate`. DO NOT EDIT.
//
// This file falls under no component and .lydite/components.yml excludes it,
// which is the half of the orphan gate `scripts/seed.ts` cannot show. One file
// proves the gate fires; it takes a second to prove an exclude clears it, and
// without that a broken exclude would look exactly like a repository that
// happened to have nothing to exclude.
//
// It is generated code, and that is the case worth carrying: lydite does not
// special-case a generated file, because recognising one means reading it and
// a gate that reads files has to be right about every language it meets. So
// the exclude is where this repository says so, in a line a reviewer sees.

export interface Counter {
  name: string;
  count: number;
}

export interface CounterList {
  counters: Counter[];
}

export async function listCounters(baseURL: string): Promise<CounterList> {
  const res = await fetch(`${baseURL}/counters`);
  if (!res.ok) {
    throw new Error(`listCounters: ${res.status}`);
  }
  return (await res.json()) as CounterList;
}

export async function getCounter(baseURL: string, name: string): Promise<Counter> {
  const res = await fetch(`${baseURL}/counters/${encodeURIComponent(name)}`);
  if (!res.ok) {
    throw new Error(`getCounter: ${res.status}`);
  }
  return (await res.json()) as Counter;
}
