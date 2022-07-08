import { getInput, setOutput } from "@actions/core";
import { simpleGit, LogResult, SimpleGit } from "simple-git";

export type Commit = LogResult["all"][0];

interface Changelog {
  module: string;
  prevTag: string | undefined;
  message?: string;
  fixes: Commit[];
  features: Commit[];
  breakingChanges: Commit[];
}

export async function generateChangelog() {
  const git = simpleGit();

  const { module, scopes, message } = getInputs();
  console.log({ module, scopes, message });

  const prevTag = await getPrevTag(git, module);
  console.log(`Previous tag: ${prevTag}`);
  let commits: LogResult;
  if (!prevTag) {
    console.log(`Getting all commits`);
    commits = await git.log();
  } else {
    console.log(`Getting commits since previous tag`);
    commits = await git.log({ from: prevTag, to: "HEAD" });
  }
  const relevantCommits = commits.all.filter(filterRelevantCommits(scopes));

  const changelog: Changelog = {
    prevTag,
    module,
    message,
    fixes: [],
    features: [],
    breakingChanges: [],
  };
  for (const commit of relevantCommits) {
    if (
      commit.message.startsWith("feat") &&
      commit.body.includes("BREAKING CHANGE")
    ) {
      changelog.breakingChanges.push(commit);
    } else if (commit.message.startsWith("feat")) {
      changelog.features.push(commit);
    } else {
      changelog.fixes.push(commit);
    }
  }

  setOutputs(changelog);
}

// Utils

function getInputs() {
  const module = getInput("module");
  const message = getInput("message") || undefined;
  const scopes = getInput("scopes")
    .split(",")
    .map((scope) => scope.trim());
  return { module, message, scopes };
}

function setOutputs(changelog: Changelog) {
  const changes =
    changelog.fixes.length +
    changelog.features.length +
    changelog.breakingChanges.length;
  if (changes === 0) {
    setOutput("changelog", "");
    setOutput("skipped", true);
    setOutput("nextVersion", "");
    return;
  }

  const lines: string[] = [];
  if (changelog.message) {
    lines.push(changelog.message);
  }
  if (changelog.breakingChanges.length > 0) {
    lines.push(
      "",
      "### BREAKING CHANGES",
      "",
      ...changelog.breakingChanges.reverse().map(formatCommit)
    );
  }
  if (changelog.features.length > 0) {
    lines.push(
      "",
      "### Features",
      "",
      ...changelog.features.reverse().map(formatCommit)
    );
  }
  if (changelog.fixes.length > 0) {
    lines.push(
      "",
      "### Fixes",
      "",
      ...changelog.fixes.reverse().map(formatCommit)
    );
  }

  const changelogText = lines.join("\n").trim();
  const nextVersion = getNextVersion(changelog);

  console.log("Changelog:");
  console.log(changelogText);
  console.log("Next version:");
  console.log(nextVersion);
  setOutput("changelog", changelogText);
  setOutput("skipped", false);
  setOutput("nextVersion", nextVersion);
}

function filterRelevantCommits(scopes: string[]) {
  const regex = new RegExp(`^(feat!?|fix)\\((${scopes.join("|")})\\)`, "m");
  return (log: Commit): boolean => {
    return regex.test(log.message);
  };
}

function getTagPrefix(module: string): string {
  return module.toLowerCase().replace(/\s/gm, "-");
}

async function getPrevTag(
  git: SimpleGit,
  module: string
): Promise<string | undefined> {
  const tags = (await git.tags()).all.reverse();
  console.log("Tags:", tags);
  return tags.find((tag) => tag.startsWith(`${getTagPrefix(module)}-v`));
}

function getNextVersion(changelog: Changelog): string {
  if (!changelog.prevTag) {
    console.log("No previous tag, using 1.0.0");
    return "1.0.0";
  }

  let [major, minor, patch] = changelog.prevTag
    .replace(`${getTagPrefix(changelog.module)}-v`, "")
    .split(".")
    .map(Number);
  console.log(`Previous version: ${major}.${minor}.${patch}`);

  if (changelog.breakingChanges.length > 0) {
    major++;
  } else if (changelog.features.length > 0) {
    minor++;
  } else if (changelog.fixes.length > 0) {
    patch++;
  }
  return `${major}.${minor}.${patch}`;
}

function formatCommit(commit: Commit): string {
  const regex = /.*\((.*?)\):\s+(.*)/;
  const match = regex.exec(commit.message);
  return `- **${match![1]}:** ${match![2]} (${commit.hash})`;
}
