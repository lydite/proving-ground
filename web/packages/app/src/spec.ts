// The app reads its route out of the OpenAPI description go/api emits. Nothing
// derives that edge from the source, so the web component declares it: a
// depends_on edge to api, and a watch on docs/openapi.json.

import openapi from "../../../../docs/openapi.json" with { type: "json" };

/** The path of the operation, taken from the spec rather than written twice. */
export function pathForOperation(operationId: string): string {
  for (const [path, methods] of Object.entries(openapi.paths)) {
    for (const operation of Object.values(methods)) {
      if (operation.operationId === operationId) {
        return path;
      }
    }
  }
  throw new Error(`no operation ${operationId} in the spec`);
}

export const countersPath = pathForOperation("listCounters");
