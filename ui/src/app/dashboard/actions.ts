"use server";

import { Timestamp } from "@bufbuild/protobuf";
import { createPromiseClient } from "@connectrpc/connect";
import { TemporalService } from "@buf/kevinmichaelchen_temporalapis.connectrpc_es/temporal/v1beta1/api_connect";
import { createConnectTransport } from "@connectrpc/connect-web";

const temporalClient = createPromiseClient(
  TemporalService,
  createConnectTransport({
    baseUrl: "http://localhost:8081",
  }),
);

export async function createOnboardingWorkflow({
  orgName,
  profileName,
  start,
  end,
}: {
  orgName: string;
  profileName: string;
  start: Date;
  end: Date;
}) {
  "use server";
  const response = await temporalClient.createOnboardingWorkflow({
    org: {
      name: orgName,
    },
    license: {
      start: Timestamp.fromDate(start),
      end: Timestamp.fromDate(end),
    },
    profile: {
      fullName: profileName,
    },
  });

  // TODO it's complaining we can't pass a RSC object to a CC object
  return response.toJson();
}
