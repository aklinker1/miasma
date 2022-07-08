import { vi, test, expect } from "vitest";
import { generateChangelog, Commit } from ".";
import { getInput, setOutput } from "@actions/core";
import { simpleGit } from "simple-git";

vi.mock("@actions/core", () => ({
  getInput: vi.fn(),
  setOutput: vi.fn(),
}));
const getInputMock = vi.mocked(getInput);
const setOutputMock = vi.mocked(setOutput);
function mockInputs(module: string, scopes: string, message?: string) {
  getInputMock.mockImplementation((name) => {
    if (name === "module") return module;
    else if (name === "scopes") return scopes;
    else if (name === "message") return message || "";
    else throw Error("Unknown input name: " + name);
  });
}

vi.mock("simple-git");
const simpleGitMock = vi.mocked(simpleGit);
function mockGit(commits: Commit[], tags: string[]) {
  const git = {
    log: vi.fn().mockResolvedValue({ all: commits }) as any,
    tags: vi.fn().mockResolvedValue({ all: tags }),
  };
  simpleGitMock.mockReturnValue(git as any);
  return git;
}

function mockCommit(override: Partial<Commit>): Commit {
  const base: Commit = {
    author_email: "random-user@example.com",
    author_name: "Random User",
    message: "fix(api): some API work 2",
    body: "Some body",
    date: Date(),
    hash: "random",
    refs: "",
  };
  return { ...base, ...override };
}

test("Example Server Changelog", async () => {
  const module = "Server";
  const scopes = "api, ui";
  const message = "Download via Docker Hub";
  mockInputs(module, scopes, message);
  const git = mockGit(
    [
      mockCommit({
        message: "fix(api): some API work 2",
        hash: "2345",
      }),
      mockCommit({
        message: "chore(api): some API chore",
        hash: "3456",
      }),
      mockCommit({
        message: "chore(release): cli-v0.5.0",
        hash: "3456",
        refs: "tag: cli-v0.5.0",
      }),
      mockCommit({
        message: "fix(cli): Some CLI work",
        hash: "4567",
      }),
      mockCommit({
        message: "feat(ui): Some UI work 2",
        body: "Some commit body",
        hash: "5678",
      }),
      mockCommit({
        message: "fix(ui): Some UI work 1",
        hash: "5678",
      }),
      mockCommit({
        message: "fix(api): Some API work 1",
        hash: "6789",
      }),
    ],
    ["cli-v0.5.0", "server-v1.1.0"]
  );

  await generateChangelog();

  expect(git.log).toBeCalledTimes(1);
  expect(git.log).toBeCalledWith({ from: "server-v1.1.0", to: "HEAD" });
  expect(setOutputMock).toBeCalledTimes(3);
  expect(setOutputMock).toBeCalledWith(
    "changelog",
    `
Download via Docker Hub

### Features

- **ui:** Some UI work 2 (5678)

### Fixes

- **api:** Some API work 1 (6789)
- **ui:** Some UI work 1 (5678)
- **api:** some API work 2 (2345)
  `.trim()
  );
  expect(setOutputMock).toBeCalledWith("skipped", false);
  expect(setOutputMock).toBeCalledWith("nextVersion", "1.2.0");
});

test("Example CLI Changelog", async () => {
  const module = "CLI";
  const scopes = "cli";
  mockInputs(module, scopes);
  const git = mockGit(
    [
      mockCommit({
        message: "feat(cli): Some CLI change 3",
        hash: "1234",
      }),
      mockCommit({
        message: "fix(cli): Some CLI change 2",
        hash: "2345",
      }),
      mockCommit({
        message: "fix(ui): Some UI change",
        hash: "2345",
      }),
      mockCommit({
        message: "feat(cli): Some CLI change 1",
        hash: "3456",
      }),
    ],
    ["server-v2.0.0"]
  );

  await generateChangelog();

  expect(git.log).toBeCalledTimes(1);
  expect(git.log).toBeCalledWith();
  expect(setOutputMock).toBeCalledTimes(3);
  expect(setOutputMock).toBeCalledWith(
    "changelog",
    `
### Features

- **cli:** Some CLI change 1 (3456)
- **cli:** Some CLI change 3 (1234)

### Fixes

- **cli:** Some CLI change 2 (2345)
    `.trim()
  );
  expect(setOutputMock).toBeCalledWith("skipped", false);
  expect(setOutputMock).toBeCalledWith("nextVersion", "1.0.0");
});
