import http from "k6/http";
import { check, sleep } from "k6";
import { Rate } from "k6/metrics";

// Custom metrics
const errorRate = new Rate("errors");

// Test configuration - adjust these based on your needs
export const options = {
  stages: [
    // Ramp up
    // { duration: "30s", target: 10 }, // Ramp up to 10 users over 30s
    // { duration: "1m", target: 50 }, // Scale up to 50 users over 1min
    // { duration: "2m", target: 100 }, // Scale up to 100 users over 2min
    // { duration: "3m", target: 200 }, // Scale up to 200 users over 3min

    // Sustained load
    // { duration: "5m", target: 200 }, // Stay at 200 users for 5min

    // Spike test
    { duration: "30s", target: 10000 }, // Spike to 500 users for 30s
    // { duration: "30s", target: 500 }, // Spike to 500 users for 30s
    { duration: "30s", target: 200 }, // Back down to 200

    // Ramp down
    // { duration: "1m", target: 0 }, // Ramp down to 0 users
  ],

  // Alternative: Simple constant load test (uncomment to use)
  // vus: 100,        // 100 virtual users
  // duration: '5m',  // for 5 minutes

  thresholds: {
    http_req_duration: ["p(95)<500"], // 95% of requests must complete below 500ms
    http_req_failed: ["rate<0.05"], // Error rate must be less than 5%
    checks: ["rate>0.9"], // 90% of checks must pass
  },
};

// Test data - add your actual bearer token here
const BASE_URL = "http://localhost:3000";
const BEARER_TOKEN =
  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiNjBhZjIxYjUtNGE5MS00Y2EzLWEwZjUtYTBiNDU4ZmY1ZTBlIiwiY3JlYXRlZF9hdCI6IjIwMjUtMDgtMTNUMDY6NDc6NDUuMDAxODk3LTA1OjAwIiwicm9sZSI6ImFkbWluIiwic3RhdHVzIjoiYWN0aXZlIn0sImV4cCI6MzMzMTg4NTY2NX0.tK0b_71Y6_pE_zc8eKRIo2XMcNYABDPLYIyrVjkmfC4"; // Replace with actual token

export default function () {
  const params = {
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${BEARER_TOKEN}`,
    },
    timeout: "30s", // 30 second timeout
  };

  // Make the request
  const response = http.get(`${BASE_URL}/api/v1/config/bedtimebully`, params);

  // Check response
  const result = check(response, {
    "status is 200": (r) => r.status === 200,
    "response time < 500ms": (r) => r.timings.duration < 500,
    "response time < 1000ms": (r) => r.timings.duration < 1000,
    "response has body": (r) => r.body && r.body.length > 0,
    "content-type is JSON": (r) =>
      r.headers["Content-Type"] &&
      r.headers["Content-Type"].includes("application/json"),
  });

  // Track errors
  errorRate.add(!result);

  // Optional: Log failed requests for debugging
  if (response.status !== 200) {
    console.error(
      `Request failed with status ${response.status}: ${response.body}`,
    );
  }

  // Small random delay between requests (0-100ms) to simulate real user behavior
  sleep(Math.random() * 0.1);
}

// Setup function (runs once before test starts)
export function setup() {
  console.log("Starting load test for bedtimebully endpoint...");
  console.log(`Target URL: ${BASE_URL}/api/v1/config/bedtimebully`);

  // Optional: Test authentication by making a single request
  const testResponse = http.get(`${BASE_URL}/api/v1/config/bedtimebully`, {
    headers: {
      Authorization: `Bearer ${BEARER_TOKEN}`,
    },
  });

  if (testResponse.status !== 200) {
    console.warn(
      `Warning: Test request failed with status ${testResponse.status}`,
    );
  }
}

// Teardown function (runs once after test completes)
export function teardown(data) {
  console.log("Load test completed!");
}
