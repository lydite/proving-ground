import { useEffect, useState } from "react";
import { CounterList, type Counter } from "@proving-ground/ui";

import { countersPath } from "./spec.js";

export function App({ baseUrl }: { baseUrl: string }) {
  const [counters, setCounters] = useState<Counter[]>([]);

  useEffect(() => {
    let live = true;
    fetch(baseUrl + countersPath)
      .then((response) => response.json())
      .then((body: Counter[]) => {
        if (live) {
          setCounters(body);
        }
      })
      .catch(() => undefined);
    return () => {
      live = false;
    };
  }, [baseUrl]);

  return <CounterList counters={counters} />;
}
