import { createRoot } from "react-dom/client";

import { App } from "./App.js";

const root = document.getElementById("root");
if (root) {
  createRoot(root).render(<App baseUrl="http://127.0.0.1:8080" />);
}
