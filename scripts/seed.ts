// Fixture rows for a local database, run by hand with `npx tsx scripts/seed.ts`.
//
// This file falls under no component, and that is what it is here to do: the
// orphan gate must fire on it. It is not part of the web workspace — `web/` is an
// npm workspace and `scripts/` is outside it — so widening the web component's
// directory is the wrong fix. Nothing tests it and no runner builds it, so it is
// not a component either. The honest resolution is an explicit exclude, and
// .lydite/components.yml deliberately does not carry one yet.

interface Row {
  name: string;
  count: number;
}

const rows: Row[] = [
  { name: "alpha", count: 1 },
  { name: "bravo", count: 0 },
  { name: "zulu", count: 3 },
];

for (const row of rows) {
  console.log(`INSERT INTO counters (name, count) VALUES ('${row.name}', ${row.count});`);
}
