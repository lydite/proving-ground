import type { Counter } from "./counters.js";
import { tally } from "./counters.js";

export function CounterList({ counters }: { counters: Counter[] }) {
  return (
    <section>
      <ul>
        {counters.map((counter) => (
          <li key={counter.name}>
            {counter.name}: {counter.count}
          </li>
        ))}
      </ul>
      <p data-testid="total">{tally(counters)}</p>
    </section>
  );
}
