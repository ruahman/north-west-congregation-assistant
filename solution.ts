const args = Deno.args;
const [cmd] = args;

const workspace = {
  "test-stencil": {
    "build": ["bun", "run", "build"],
  },
  "test-solidjs": {
    "build": ["bun", "run", "build"],
  },
  "test-astro": {
    "build": ["bun", "run", "build"],
    "run": {
      "cmd": ["bun", "run", "dev"],
      "message": "test-astro: checkout http://localhost:4321/",
    },
  },
  "test-deno": {
    "run": {
      "cmd": ["deno", "run", "-A", "main.ts"],
      "message": "test-deno: checkout http://localhost:8000/",
    },
  },
};

async function build() {
  for (const [proj, options] of Object.entries(workspace)) {
    if ("build" in options) {
      const [cmd, ...args] = options["build"];

      const command = new Deno.Command(cmd, {
        args: args,
        cwd: `./${proj}`,
      });
      const { code, stdout, stderr } = await command.output();
      console.log(`code: ${code}`);
      console.log(`stdout: ${new TextDecoder().decode(stdout)}`);
      console.log(`stderr: ${new TextDecoder().decode(stderr)}`);
    }
  }
}

async function run() {
  const promises = [];
  for (const [proj, options] of Object.entries(workspace)) {
    if ("run" in options) {
      const [cmd, ...args] = options["run"]["cmd"];
      const command = new Deno.Command(cmd, {
        args: args,
        cwd: `./${proj}`,
      });
      const message = options["run"]["message"];
      console.log(message);
      promises.push(command.output());
    }
  }
  await Promise.all(promises);
}

function test() {
  console.log("testing...");
}

function link() {
  console.log("linking...");
}

if (cmd == "build") {
  await build();
} else if (cmd == "run") {
  await run();
} else if (cmd == "test") {
  test();
} else if (cmd == "link") {
  link();
}
