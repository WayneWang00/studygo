# Roo-Code-Docs

## README

# Roo Code Docs

This website is built using [Docusaurus](https://docusaurus.io/), a modern static website generator, and lives at https://docs.roocode.com

### Installation

```
$ npm install
```

### Local Development

```
$ npm start
```

This command starts a local development server and opens up a browser window. Most changes are reflected live without having to restart the server.

### License

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)

## tutorial-videos

---
description: Watch comprehensive video tutorials to learn how to use Roo Code effectively. From basics to advanced features, master AI-assisted development.
keywords:
  - Roo Code tutorials
  - video tutorials
  - AI coding tutorials
  - learn Roo Code
  - video guides
image: /img/social-share.jpg
---

import VideoGrid from '@site/src/components/VideoGrid';

# Tutorial Videos

Learn how to build powerful applications and enhance your development workflow with these hands-on Roo Code tutorials. Big thanks to all of the creators!

<VideoGrid />

## tips-and-tricks

---
description: Discover expert tips and tricks for using Roo Code effectively. Learn best practices, productivity hacks, and advanced techniques.
keywords:
  - Roo Code tips
  - productivity tips
  - best practices
  - AI coding tips
  - advanced techniques
image: /img/social-share.jpg
---

# Tips & Tricks

A collection of quick tips to help you get the most out of Roo Code.

- Drag Roo Code to the [Secondary Sidebar](https://code.visualstudio.com/api/ux-guidelines/sidebars#secondary-sidebar) so you can see the Explorer, Search, Source Control, etc.
<img src="/img/right-column-roo.gif" alt="Put Roo on the Right Column" width="900" />
- Once you have Roo Code in a separate sidebar from the file explorer, you can drag files from the explorer into the chat window (and even multiple at once). Just make sure to hold down the shift key after you start dragging the files.
- If you're not using [MCP](/features/mcp/overview), turn it off in the <Codicon name="server" /> MCP Servers tab to significantly cut down the size of the system prompt.
- To keep your [custom modes](/features/custom-modes) on track, limit the types of files that they're allowed to edit.
- If you hit the dreaded `input length and max tokens exceed context limit` error, you can recover by deleting a message, rolling back to a previous checkpoint, or switching over to a model with a long context window like Gemini for a message.
- In general, be thoughtful about your `Max Tokens` setting for thinking models. Every token you allocate to that takes away from space available to store conversation history. Consider only using high `Max Tokens` / `Max Thinking Tokens` settings with modes like Architect and Debug, and keeping Code mode at 16k max tokens or less.
- If there's a real world job posting for something you want a custom mode to do, try asking Code mode to `Create a custom mode based on the job posting at @[url]`
- If you want to really accelerate, check out multiple copies of your repository and run Roo Code on all of them in parallel (using git to resolve any conflicts, same as with human devs).
- When using Debug mode, ask Roo to "start a new task in Debug mode with all of the necessary context needed to figure out X" so that the debugging process uses its own context window and doesn't pollute the main task
- Add your own tips by clicking "Edit this page" below!
- To manage large files and reduce context/resource usage, adjust the `File read auto-truncate threshold` setting. This setting controls the number of lines read from a file in one batch. Lower values can improve performance when working with very large files, but may require more read operations. You can find this setting in the Roo Code settings under 'Advanced Settings'.
- Set up a keyboard shortcut for the [`roo.acceptInput` command](/features/keyboard-shortcuts) to accept suggestions or submit text input without using the mouse. Perfect for keyboard-focused workflows and reducing hand strain.
- Use **Sticky Models** to assign specialized AI models to different modes (reasoning model for planning, non-reasoning model for coding). Roo automatically switches to each mode's last-used model without manual selection.
- Customize the [context reduction prompt](/features/intelligent-context-condensing#customizing-the-context-condensing-prompt) if you find that for your domain/use case it forgets particular things. You can instruct it to preserve specific types of information that are critical to your workflow.

## index

---
sidebar_label: Welcome
description: Learn how to use Roo Code, the autonomous AI coding agent that helps you code faster and smarter in VS Code. Complete guides, tutorials, and documentation.
keywords:
  - Roo Code
  - AI coding assistant
  - VS Code extension
  - autonomous coding agent
  - AI pair programmer
  - code generation
  - documentation
image: /img/social-share.jpg
---

# Roo Code Docs

Roo Code is an AI-powered autonomous coding agent that lives in your editor. It helps you code faster and smarter, whether you're starting a new project, maintaining existing code, or learning new technologies.

---

## What Can Roo Code Do?

- **Generate Code** from natural language descriptions
- **Refactor & Debug** existing code
- **Write & Update** documentation
- **Answer Questions** about your codebase
- **Automate** repetitive tasks
- **Create** new files and projects

---

## Quick Start

1. [Install Roo Code](/getting-started/installing)
2. [Connect Your AI Provider](/getting-started/connecting-api-provider)
3. [Try Your First Task](/getting-started/your-first-task)

---

## Key Features

### Multiple Modes
Roo Code adapts to your needs with specialized [modes](/basic-usage/using-modes):
- **Code Mode:** For general-purpose coding tasks
- **Architect Mode:** For planning and technical leadership
- **Ask Mode:** For answering questions and providing information
- **Debug Mode:** For systematic problem diagnosis
- **Orchestrator Mode:** For managing complex tasks and delegating work
- **[Custom Modes](/features/custom-modes):** Create unlimited specialized personas for security auditing, performance optimization, documentation, or any other task

### Smart Tools
Roo Code comes with powerful [tools](/basic-usage/how-tools-work) that can:
- Read and write files in your project
- Execute commands in your VS Code terminal
- Control a web browser
- Use external tools via [MCP (Model Context Protocol)](/features/mcp/overview)

MCP extends Roo Code's capabilities by allowing you to add unlimited custom tools. Integrate with external APIs, connect to databases, or create specialized development tools - MCP provides the framework to expand Roo Code's functionality to meet your specific needs.

### Customization
Make Roo Code work your way with:
- [Custom Instructions](/features/custom-instructions) for personalized behavior
- [Custom Modes](/features/custom-modes) for specialized tasks
- [Local Models](/advanced-usage/local-models) for offline use
- [Auto-Approval Settings](/features/auto-approving-actions) for faster workflows

---

## Resources

### Documentation
- [Basic Usage Guide](/basic-usage/the-chat-interface)
- [Advanced Features](/features/auto-approving-actions)
- [Frequently Asked Questions](/faq)

### Community & Socials
- **Discord:** [Join our Discord server](https://discord.gg/roocode) for real-time help and discussions.
- **Reddit:** [Visit our subreddit](https://www.reddit.com/r/RooCode) to share experiences and tips.
- **GitHub:** Report [issues](https://github.com/RooCodeInc/Roo-Code/issues) or request [features](https://github.com/RooCodeInc/Roo-Code/discussions/categories/feature-requests?discussions_q=is%3Aopen+category%3A%22Feature+Requests%22+sort%3Atop).
- **X (Twitter):** [Follow @roo_code](https://x.com/roo_code).
- **Bluesky:** [Follow roocode.bsky.social](https://bsky.app/profile/roocode.bsky.social).
- **LinkedIn:** [Follow Roo Code](https://www.linkedin.com/company/roo-code).

Ready to get started? Click the **Next** button below to begin your journey with Roo Code!

## faq

---
description: Find answers to common questions about Roo Code, including setup, usage, troubleshooting, and advanced features. Get help with API keys, modes, and more.
keywords:
  - Roo Code FAQ
  - frequently asked questions
  - troubleshooting
  - API setup
  - custom modes
  - MCP
  - local models
image: /img/social-share.jpg
---

import KangarooIcon from '@site/src/components/KangarooIcon';

# Frequently Asked Questions

This page answers some common questions about Roo Code.

---

## General

### What is Roo Code?

Roo Code is an AI-powered autonomous coding agent that lives in your editor.

### How does Roo Code work?

Roo Code uses large language models (LLMs) to understand your requests and translate them into actions.  It can:

*   Read and write files in your project.
*   Execute commands in your VS Code terminal.
*   Perform web browsing (if enabled).
*   Use external tools via the Model Context Protocol (MCP).

You interact with Roo Code through a chat interface, where you provide instructions and review/approve its proposed actions.

### What can Roo Code do?

Roo Code can help with a variety of coding tasks, including:

*   Generating code from natural language descriptions.
*   Refactoring existing code.
*   Fixing bugs.
*   Writing documentation.
*   Explaining code.
*   Answering questions about your codebase.
*   Automating repetitive tasks.
*   Creating new files and projects.

### Is Roo Code free to use?

The Roo Code extension itself is free and open-source. However, Roo Code relies on external API providers (like [Anthropic](providers/anthropic), [OpenAI](providers/openai), [OpenRouter](providers/openrouter), [Requesty](providers/requesty), etc.) for its AI capabilities.  These providers typically charge for API usage based on the number of tokens processed.  You will need to create an account and obtain an API key from your chosen provider.  See [Setting Up Your First AI Provider](getting-started/connecting-api-provider) for details.

### What are the risks of using Roo Code?

Roo Code is a powerful tool, and it's important to use it responsibly.  Here are some things to keep in mind:

*   **Roo Code can make mistakes.**  Always review Roo Code's proposed changes carefully before approving them.
*   **Roo Code can execute commands.**  Be very cautious about allowing Roo Code to run commands, especially if you're using auto-approval.
*   **Roo Code can access the internet.** If you're using a provider that supports web browsing, be aware that Roo Code could potentially access sensitive information.

---

## Setup & Installation

### How do I install Roo Code?

See the [Installation Guide](/getting-started/installing) for detailed instructions.

### Which API providers are supported?

Roo Code supports a wide range of API providers, including:
*   [Anthropic (Claude)](/providers/anthropic)
*   [OpenAI](/providers/openai)
*   [OpenRouter](/providers/openrouter)
*   [Google Gemini](/providers/gemini)
*   [Glama](/providers/glama)
*   [AWS Bedrock](/providers/bedrock)
*   [GCP Vertex AI](/providers/vertex)
*   [Ollama](/providers/ollama)
*   [LM Studio](/providers/lmstudio)
*   [DeepSeek](/providers/deepseek)
*   [Mistral](/providers/mistral)
*   [Unbound](/providers/unbound)
*   [Requesty](/providers/requesty)
*   [VS Code Language Model API](/providers/vscode-lm)

### How do I get an API key?
Each API provider has its own process for obtaining an API key.  See the [Setting Up Your First AI Provider](/getting-started/connecting-api-provider) for links to the relevant documentation for each provider.

### Can I use Roo Code with local models?
Yes, Roo Code supports running models locally using [Ollama](/providers/ollama) and [LM Studio](/providers/lmstudio).  See [Using Local Models](/advanced-usage/local-models) for instructions.

---

## Usage

### How do I start a new task?
Open the Roo Code panel (<KangarooIcon />) and type your task in the chat box. Be clear and specific about what you want Roo Code to do. See [Typing Your Requests](/basic-usage/typing-your-requests) for best practices.

### What are modes in Roo Code?
[Modes](/basic-usage/using-modes) are different personas that Roo Code can adopt, each with a specific focus and set of capabilities. The built-in modes are:

*   **Code:** For general-purpose coding tasks.
*   **Architect:** For planning and technical leadership.
*   **Ask:** For answering questions and providing information.
*   **Debug:** For systematic problem diagnosis.
You can also create [Custom Modes](/features/custom-modes).

### How do I switch between modes?

Use the dropdown menu in the chat input area to select a different mode, or use the `/` command to switch to a specific mode.

### What are tools and how do I use them?
[Tools](/basic-usage/how-tools-work) are how Roo Code interacts with your system.  Roo Code automatically selects and uses the appropriate tools to complete your tasks. You don't need to call tools directly. You will be prompted to approve or reject each tool use.

### What are context mentions?
[Context mentions](/basic-usage/context-mentions) are a way to provide Roo Code with specific information about your project, such as files, folders, or problems. Use the "@" symbol followed by the item you want to mention (e.g., `@/src/file.ts`, `@problems`).

### Can Roo Code access the internet?

Yes, if you are using a provider with a model that support web browsing. Be mindful of the security implications of allowing this.

### Can Roo Code run commands in my terminal?

Yes, Roo Code can execute commands in your VS Code terminal.  You will be prompted to approve each command before it's executed, unless you've enabled auto-approval for commands. Be extremely cautious about auto-approving commands. If you're experiencing issues with terminal commands, see the [Shell Integration Guide](/features/shell-integration) for troubleshooting.

### How do I provide feedback to Roo Code?

You can provide feedback by approving or rejecting Roo Code's proposed actions. You can provide additional feedback by using the feedback field.

### Can I customize Roo Code's behavior?

Yes, you can customize Roo Code in several ways:

*   **Custom Instructions:** Provide general instructions that apply to all modes, or mode-specific instructions.
*   **Custom Modes:** Create your own modes with tailored prompts and some tool permissions.
*   **`.roorules` Files:** Create `.roorules` files in your project to provide additional guidelines.
*   **Settings:** Adjust various settings, such as auto-approval, diff editing, and more.

### Does Roo Code have any auto approval settings?
Yes, Roo Code has a few settings that when enabled will automatically approve actions. Find out more [here](/features/auto-approving-actions).

---

## Advanced Features

### Can I use Roo offline?
Yes, if you use a [local model](/advanced-usage/local-models).

### What is MCP (Model Context Protocol)?
[MCP](/features/mcp/overview) is a protocol that allows Roo Code to communicate with external servers, extending its capabilities with custom tools and resources.

### Can I create my own MCP servers?

Yes, you can create your own MCP servers to add custom functionality to Roo Code. See the [MCP documentation](https://github.com/modelcontextprotocol) for details.

### What is Codebase Indexing?

[Codebase Indexing](/features/codebase-indexing) creates a semantic search index of your project using AI embeddings. This enables Roo Code to better understand and navigate large codebases by finding relevant code based on meaning rather than just keywords.

### How much does Codebase Indexing cost?

Codebase Indexing requires an OpenAI API key for generating embeddings and a Qdrant vector database for storage. Costs depend on your project size and the embedding model used. Initial indexing is the most expensive part; subsequent updates are incremental and much cheaper.

---

## Troubleshooting

### Roo Code isn't responding. What should I do?

*   Make sure your API key is correct and hasn't expired.
*   Check your internet connection.
*   Check the status of your chosen API provider.
*   Try restarting VS Code.
*   If the problem persists, report the issue on [GitHub](https://github.com/RooCodeInc/Roo-Code/issues) or [Discord](https://discord.gg/roocode).

### I'm seeing an error message. What does it mean?

The error message should provide some information about the problem. If you're unsure how to resolve it, seek help in [Discord](https://discord.gg/roocode).

### Roo Code made changes I didn't want. How do I undo them?

Roo Code uses VS Code's built-in file editing capabilities.  You can use the standard "Undo" command (Ctrl/Cmd + Z) to revert changes. Also, if experimental checkpoints are enabled, Roo can revert changes made to a file.

### Roo Code can't write to markdown files. What's wrong?

If Roo Code fails to write to `.md` files with errors like "Failed to open diff editor" or "write_to_file tool failed", this is typically caused by VS Code extensions or settings that interfere with file editing:

**Common causes:**
- Extensions with "format on save" functionality
- VS Code settings that open markdown files in preview mode by default
- The Markdown Preview extension or similar markdown processing extensions

**Solutions:**
- Disable any extensions that automatically format files on save
- Remove these settings from your VS Code `settings.json`:
  ```json
  "markdown.preview.openMarkdownLinks": "inPreview",
  "workbench.editorAssociations": {
    "*.md": "vscode.markdown.preview.editor"
  }
  ```
- Temporarily disable markdown-related extensions to test if they're causing the issue
- Restart VS Code after making these changes

### How do I report a bug or suggest a feature?

Please report bugs or suggest features on the Roo Code [Issues page](https://github.com/RooCodeInc/Roo-Code/issues) and [Feature Requests page](https://github.com/RooCodeInc/Roo-Code/discussions/categories/feature-requests?discussions_q=is%3Aopen+category%3A%22Feature+Requests%22+sort%3Atop).

## context-poisoning

---
description: Learn about context poisoning in AI coding assistants, its symptoms, causes, and effective recovery strategies to maintain accurate AI responses.
keywords:
  - context poisoning
  - AI accuracy
  - Roo Code troubleshooting
  - LLM context management
  - session recovery
image: /img/social-share.jpg
---

# Context Poisoning

:::info
Context poisoning is a persistent issue within a given session. Once a chat session's context is compromised, treat that session as disposable. Starting fresh with a clean context is crucial for maintaining the accuracy and effectiveness of your Roo Code agent.
:::

Context poisoning occurs when inaccurate or irrelevant data contaminates the language model's active context. This leads the model to draw incorrect conclusions, provide erroneous information to tools, and progressively deviate from the intended task with each interaction.

---

## Symptoms of Context Poisoning

Identify context poisoning by observing these behaviors:

*   **Degraded Output Quality:** Suggestions become nonsensical, repetitive, or irrelevant.
*   **Tool Misalignment:** Tool calls no longer correspond to the user's requests.
*   **Orchestration Failures:** Orchestrator chains may stall, loop indefinitely, or fail to complete.
*   **Temporary Fixes:** Re-applying a clean prompt or instructions offers only brief respite before issues resurface.
*   **Tool Usage Confusion:** The model struggles to correctly use or recall how to use tools defined in the system prompt.

---

## Common Causes

Context poisoning can be triggered by several factors:

*   **Model Hallucination:** The model generates an incorrect piece of information and subsequently treats it as a factual part of the context.
*   **Code Comments:** Outdated, incorrect, or ambiguous comments in the codebase can be misinterpreted by the model, leading it down the wrong path.
*   **Contaminated User Input:** Copy-pasting logs or text containing hidden or rogue control characters.
*   **Context Window Overflow:** As a session grows, older, useful information may be pushed out of the model's limited context window, allowing "poisoned" data to have a greater relative impact.

Once bad data enters the context, it tends to persist. The model re-evaluates this tainted information in subsequent reasoning cycles, similar to a permanent flaw affecting its perception until the context is completely reset.

---

## Can a "Wake-Up Prompt" Resolve Context Poisoning?

**Short Answer:** No.

A corrective prompt might temporarily suppress symptoms, but the problematic data remains in the conversational buffer. The model will likely revert to the poisoned state as soon as the interaction deviates from the narrow scope of the corrective prompt.

**Detailed Explanation:**

*   Re-injecting the full set of tool definitions or core directives can sometimes mask the damage for one or some interactions following the initial context poisoning .
*   However, the underlying poisoned context remains. Any query or task outside the immediate "patch" will likely re-trigger the original issue.
*   This approach is unreliable, akin to placing a warning label on a leaking pipe instead of repairing it.

---

## Effective Recovery Strategies

To reliably recover from context poisoning:

*   **Hard Reset the Session:** The most dependable solution is to start a new chat session. This clears the contaminated context entirely.
*   **Minimize Manual Data Dumps:** When pasting logs or other data, be selective. Only include the essential information the model requires.
*   **Manage Context Window Size:** For large or complex tasks, consider breaking them into smaller, focused chat sessions. This helps ensure that stale or irrelevant information ages out of the context window more quickly.
*   **Validate Tool Output:** If a tool returns nonsensical or clearly incorrect data, delete that message from the chat history before the model can process it and incorporate it into its context.

---

## Addressing a Common Question: The "Magic Bullet" Prompt

A frequent question from the community is:
> "Have you found a prompt that wakes it back up? Maybe a prompt that just has the tools instructions we can push back in manually?”

As explained, no single prompt offers a lasting fix. Any immediate improvement is superficial because the corrupted lines of text persist in the session's history, ready to cause further issues. The only robust solution is to discard the compromised session, initiate a new one, and provide it with a clean prompt and the correct tool definitions from the outset.

## **Footgun Prompting**

---
sidebar_label: 'Footgun Prompting'
description: Advanced guide to overriding system prompts in Roo Code. Learn how to customize AI behavior with caution using system prompt overrides.
keywords:
  - footgun prompting
  - system prompt override
  - advanced customization
  - Roo Code configuration
  - custom prompts
image: /img/social-share.jpg
---

# Footgun Prompting: Override System Prompts

Footgun Prompting (System Prompt Override) lets you replace the default system prompt for a specific Roo Code mode. This offers granular control but bypasses built-in safeguards. Use with caution.

<img src="/img/footgun-prompting/footgun-prompting-1.png" alt="Warning indicator for active system prompt override" width="600" />
**Warning Indicator:** When a system prompt override is active for the current mode, Roo Code will display a warning icon in the chat input area to remind you that the default behavior has been modified.


:::info **footgun** _(noun)_

1.  _(programming slang, humorous, derogatory)_ Any feature likely to lead to the programmer shooting themself in the foot.

> The System Prompt Override is considered a footgun because modifying the core instructions without a deep understanding can lead to unexpected or broken behavior, especially regarding tool usage and response consistency.

:::

---

## How It Works

1.  **Override File:** Create a file named `.roo/system-prompt-{mode-slug}` in your workspace root (e.g., `.roo/system-prompt-code` for the Code mode).
2.  **Content:** The content of this file becomes the new system prompt for that specific mode.
3.  **Activation:** Roo Code automatically detects this file. When present, it replaces most of the standard system prompt sections.
4.  **Preserved Sections:** Only the core `roleDefinition` and any `customInstructions` you've set for the mode are kept alongside your override content. Standard sections like tool descriptions, rules, and capabilities are bypassed.
5.  **Construction:** The final prompt sent to the model looks like this:

    ```
    ${roleDefinition}
    
    ${content_of_your_override_file}
    
    ${customInstructions}
    ```

---

## Accessing the Feature

Find the option and instructions in the Roo Code UI:

1.  Click the <Codicon name="notebook" /> icon in the Roo Code top menu bar.
2.  Expand the **"Advanced: Override System Prompt"** section.
3.  Clicking the file path link within the explanation will open or create the correct override file for the currently selected mode in VS Code.

<img src="/img/footgun-prompting/footgun-prompting.png" alt="UI showing the Advanced: Override System Prompt section" width="500" />

---

## Using Context Variables

When creating your custom system prompt file, you can use special variables (placeholders) that Roo Code will automatically replace with relevant information about the current environment. This allows you to make your prompts more dynamic and context-aware.

Here are the available variables:

- `{{mode}}`: The slug (short name) of the current Roo Code mode being used (e.g., `code`, `chat-mode`).
- `{{language}}`: The display language configured in VS Code (e.g., `en`, `es`).
- `{{shell}}`: The default terminal shell configured in VS Code (e.g., `/bin/bash`, `powershell.exe`).
- `{{operatingSystem}}`: The type of operating system your computer is running (e.g., `Linux`, `Darwin` for macOS, `Windows_NT`).
- `{{workspace}}`: The file path to the root of your current project workspace.

**Example Usage:**

You can include these variables directly in your prompt file content like this:

```
You are assisting a user in the '{{mode}}' mode.
Their operating system is {{operatingSystem}} and their default shell is {{shell}}.
The project is located at: {{workspace}}.
Please respond in {{language}}.
```

Roo Code substitutes these placeholders before sending the prompt to the model.

---

## Key Considerations & Warnings

- **Intended Audience:** Best suited for users deeply familiar with Roo Code's prompting system and the implications of modifying core instructions.
- **Impact on Functionality:** Custom prompts override standard instructions, including those for tool usage and response consistency. This can cause unexpected behavior or errors if not managed carefully.
- **Mode-Specific:** Each override file applies only to the mode specified in its filename (`{mode-slug}`).
- **No File, No Override:** If the `.roo/system-prompt-{mode-slug}` file doesn't exist, Roo Code uses the standard system prompt generation process for that mode.
- **Blank Files Ignored:** If the override file exists but is empty (blank), it will be ignored and the default system prompt will be used.
- **Directory Creation:** Roo Code ensures the `.roo` directory exists before attempting to read or create the override file.
Use this feature cautiously. Incorrect implementation can significantly degrade Roo Code's performance and reliability for the affected mode.

## **Large Projects**

---
description: Learn strategies for effectively using Roo Code with large codebases. Manage context limits, optimize token usage, and handle complex refactoring tasks.
keywords:
  - large projects
  - context management
  - token optimization
  - codebase refactoring
  - Roo Code scalability
image: /img/social-share.jpg
---

# Working with Large Projects

Roo Code can be used with projects of any size, but large projects require some extra care to manage context effectively. Here are some tips for working with large codebases:

---

## Understanding Context Limits

Roo Code uses large language models (LLMs) that have a limited "context window."  This is the maximum amount of text (measured in tokens) that the model can process at once.  If the context is too large, the model may not be able to understand your request or generate accurate responses.

The context window includes:

*   The system prompt (instructions for Roo Code).
*   The conversation history.
*   The content of any files you mention using `@`.
*   The output of any commands or tools Roo Code uses.

---

## Strategies for Managing Context

1.  **Be Specific:**  When referring to files or code, use specific file paths and function names.  Avoid vague references like "the main file."

2.  **Use Context Mentions Effectively:** Use `@/path/to/file.ts` to include specific files. Use `@problems` to include current errors and warnings.  Use `@` followed by a commit hash to reference specific Git commits.

3.  **Break Down Tasks:** Divide large tasks into smaller, more manageable sub-tasks.  This helps keep the context focused.

4.  **Summarize:**  If you need to refer to a large amount of code, consider summarizing the relevant parts in your prompt instead of including the entire code.

5.  **Prioritize Recent History:** Roo Code automatically truncates older messages in the conversation history to stay within the context window. Be mindful of this, and re-include important context if needed.

6.  **Use Prompt Caching (if available):** Some API providers like Anthropic, OpenAI, OpenRouter and Requesty support "prompt caching". This caches your prompts for use in future tasks and helps reduce the cost and latency of requests.

---

## Example: Refactoring a Large File

Let's say you need to refactor a large TypeScript file (`src/components/MyComponent.tsx`).  Here's a possible approach:

1.  **Initial Overview:**
    ```
    @/src/components/MyComponent.tsx List the functions and classes in this file.
    ```

2.  **Target Specific Functions:**
    ```
    @/src/components/MyComponent.tsx Refactor the `processData` function to use `async/await` instead of Promises.
    ```

3.  **Iterative Changes:**  Make small, incremental changes, reviewing and approving each step.

By breaking down the task and providing specific context, you can work effectively with large files even with a limited context window.

## **Local Setup**

---
description: Learn how to set up Roo Code for local development, build from source, and contribute to the project. Complete developer setup guide.
keywords:
  - local development
  - build from source
  - Roo Code development
  - contribute
  - developer setup
image: /img/social-share.jpg
sidebar_label: Local Development Setup
---

## Local Setup & Development

:::note Developer Information Only
This section is for developers contributing to Roo Code or those who wish to build and run the extension from source.
:::

To set up Roo Code for local development:

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/RooCodeInc/Roo-Code.git
    cd Roo-Code
    ```

2.  **Install dependencies:**
    Roo Code uses `pnpm` for package management.
    ```bash
    pnpm install
    ```

3.  **Run the extension for development:**
    You have two primary methods:

    *   **Directly in VS Code (Recommended for active development):**
        *   Open the cloned `Roo-Code` folder in VS Code.
        *   Press `F5` (or select **Run** → **Start Debugging** from the menu). This opens a new VS Code window (Extension Development Host) with Roo Code running.
        *   Changes to the webview code (under `webview-ui/`) will often update live.
        *   Changes to the core extension code (under `src/`) typically require restarting the Extension Development Host window.

    *   **Building and Installing a VSIX package:**
        This method is useful for testing the packaged extension or distributing development builds.
        1.  Build the VSIX file:
            ```bash
            pnpm vsix
            ```
            A `.vsix` file (e.g., `roo-cline-<version>.vsix`) will be created in the `bin/` directory.

        2.  Install the VSIX in VS Code:
            *   **Using the command line:**
                ```bash
                code --install-extension bin/roo-cline-<version>.vsix
                ```
                (Replace `<version>` with the actual version number in the filename).
            *   **Using the VS Code interface:**
                *   Open the Extensions view (`Ctrl+Shift+X` or `Cmd+Shift+X`).
                *   Click the "..." (More Actions) menu in the Extensions view sidebar.
                *   Select "Install from VSIX..."
                *   Browse to and select the generated `.vsix` file from the `bin/` directory.

                <img src="/img/installing/installing-2.png" alt="VS Code's Install from VSIX dialog" width="600" />
                *Installing a development VSIX using VS Code's "Install from VSIX..." dialog.*

## **Local Models**

---
description: Learn how to run Roo Code with local AI models using Ollama and LM Studio. Complete setup guide for offline AI coding assistance.
keywords:
  - local models
  - Ollama
  - LM Studio
  - offline AI
  - local LLM
  - self-hosted AI
  - privacy-focused AI
image: /img/social-share.jpg
---

# Using Local Models

Roo Code supports running language models locally on your own machine using [Ollama](https://ollama.com/) and [LM Studio](https://lmstudio.ai/).  This offers several advantages:

*   **Privacy:** Your code and data never leave your computer.
*   **Offline Access:**  You can use Roo Code even without an internet connection.
*   **Cost Savings:**  Avoid API usage fees associated with cloud-based models.
*   **Customization:**  Experiment with different models and configurations.

**However, using local models also has some drawbacks:**

*   **Resource Requirements:**  Local models can be resource-intensive, requiring a powerful computer with a good CPU and, ideally, a dedicated GPU.
*   **Setup Complexity:**  Setting up local models can be more complex than using cloud-based APIs.
*   **Model Performance:**  The performance of local models can vary significantly. While some are excellent, they may not always match the capabilities of the largest, most advanced cloud models.
* **Limited Features**: Local models (and many online models) often do not support advanced features such as prompt caching, computer use, and others.

---

## Supported Local Model Providers

Roo Code currently supports two main local model providers:

1.  **Ollama:**  A popular open-source tool for running large language models locally.  It supports a wide range of models.
2.  **LM Studio:**  A user-friendly desktop application that simplifies the process of downloading, configuring, and running local models.  It also provides a local server that emulates the OpenAI API.

---

## Setting Up Local Models

For detailed setup instructions, see:
* [Setting up Ollama](/providers/ollama)
* [Setting up LM Studio](/providers/lmstudio)

Both providers offer similar capabilities but with different user interfaces and workflows. Ollama provides more control through its command-line interface, while LM Studio offers a more user-friendly graphical interface.

---

## Troubleshooting

*   **"No connection could be made because the target machine actively refused it":**  This usually means that the Ollama or LM Studio server isn't running, or is running on a different port/address than Roo Code is configured to use.  Double-check the Base URL setting.
*   **Slow Response Times:** Local models can be slower than cloud-based models, especially on less powerful hardware.  If performance is an issue, try using a smaller model.
*   **Model Not Found:** Ensure you have typed in the name of the model correctly. If you're using Ollama, use the same name that you provide in the `ollama run` command.

## **Prompt Engineering**

---
description: Master the art of writing effective prompts for Roo Code. Learn principles, techniques, and examples to get better AI coding assistance results.
keywords:
  - prompt engineering
  - AI prompts
  - effective communication
  - Roo Code tips
  - custom instructions
image: /img/social-share.jpg
---

# Prompt Engineering Tips

Prompt engineering is the art of crafting effective instructions for AI models like Roo Code.  Well-written prompts lead to better results, fewer errors, and a more efficient workflow.

---

## General Principles

*   **Be Clear and Specific:** Clearly state what you want Roo Code to do. Avoid ambiguity.
    *   **Bad:** Fix the code.
    *   **Good:** Fix the bug in the `calculateTotal` function that causes it to return incorrect results.

*   **Provide Context:** Use [Context Mentions](/basic-usage/context-mentions) to refer to specific files, folders, or problems.
    *   **Good:** `@/src/utils.ts` Refactor the `calculateTotal` function to use async/await.

*   **Break Down Tasks:** Divide complex tasks into smaller, well-defined steps.

*   **Give Examples:** If you have a specific coding style or pattern in mind, provide examples.

*   **Specify Output Format:** If you need the output in a particular format (e.g., JSON, Markdown), specify it in the prompt.

*   **Iterate:** Don't be afraid to refine your prompt if the initial results aren't what you expect.

---

## Thinking vs. Doing

It's often helpful to guide Roo Code through a "think-then-do" process:

1.  **Analyze:** Ask Roo Code to analyze the current code, identify problems, or plan the approach.
2.  **Plan:**  Have Roo Code outline the steps it will take to complete the task.
3.  **Execute:**  Instruct Roo Code to implement the plan, one step at a time.
4.  **Review:**  Carefully review the results of each step before proceeding.

---

## Using Custom Instructions

You can provide custom instructions to further tailor Roo Code's behavior. There are two types of custom instructions:

*   **Global Custom Instructions:** Apply to all modes.
*   **Mode-Specific Custom Instructions:** Apply only to a specific mode (e.g., Code, Architect, Ask, Debug, or a custom mode).

Custom instructions are added to the system prompt, providing persistent guidance to the AI model. You can use these to:

*   Enforce coding style guidelines.
*   Specify preferred libraries or frameworks.
*   Define project-specific conventions.
*   Adjust Roo Code's tone or personality.

See the [Custom Instructions](/features/custom-instructions) section for more details.

---

## Handling Ambiguity

If your request is ambiguous or lacks sufficient detail, Roo Code might:

*   **Make Assumptions:**  It might proceed based on its best guess, which may not be what you intended.
*   **Ask Follow-Up Questions:** It might use the `ask_followup_question` tool to clarify your request.

It's generally better to provide clear and specific instructions from the start to avoid unnecessary back-and-forth.

---

## Providing Feedback

If Roo Code doesn't produce the desired results, you can provide feedback by:

*   **Rejecting Actions:** Click the "Reject" button when Roo Code proposes an action you don't want.
*   **Providing Explanations:** When rejecting, explain *why* you're rejecting the action.  This helps Roo Code learn from its mistakes.
*   **Rewording Your Request:** Try rephrasing your initial task or providing more specific instructions.
*   **Manually Correcting:** If there are a few small issues, you can also directly modify the code before accepting the changes.

---

## Examples

**Good Prompt:**

> `@/src/components/Button.tsx` Refactor the `Button` component to use the `useState` hook instead of the `useReducer` hook.

**Bad Prompt:**

> Fix the button.

**Good Prompt:**

> Create a new file named `utils.py` and add a function called `calculate_average` that takes a list of numbers and returns their average.

**Bad Prompt:**

> Write some Python code.

**Good Prompt:**

> `@problems` Address all errors and warnings in the current file.

**Bad Prompt:**

> Fix everything.

By following these tips, you can write effective prompts that get the most out of Roo Code's capabilities.

## **Prompt Structure**

---
description: Understand the technical structure of prompts in Roo Code. Learn how messages are constructed, system prompts work, and optimize your interactions.
keywords:
  - prompt structure
  - system prompt
  - message flow
  - technical documentation
  - LLM communication
image: /img/social-share.jpg
---

# Prompt Structure

This page explains the technical structure of prompts in Roo Code - how messages are constructed and sent to the Large Language Model (LLM).

---

## Core Message Types

Roo Code uses three primary message types when communicating with LLMs:

- **System Prompt**: The initial instructions that define Roo's capabilities, persona, and operational rules
- **User Messages**: Content sent by you (the user) to Roo
- **Assistant Messages**: Responses generated by the LLM based on your requests

At the API level, there's also a fourth message role:

- **Tool Messages**: Results returned from tool executions, sent back to the LLM as input

Understanding these message types helps you work more effectively with Roo and can be valuable for troubleshooting or advanced customization.

---

## System Prompt

The system prompt is the foundation of Roo's behavior. It contains:

- **Role Definition**: The core persona instructions based on the selected mode (Code, Ask, Debug, etc.)
- **Tool Descriptions**: Detailed information about available tools, including parameters and examples
- **Tool Use Guidelines**: Rules for how tools should be used (sequential execution, waiting for results)
- **Capabilities**: Description of what Roo can do in the current environment
- **Available Modes**: List of all available modes and their descriptions
- **Operational Rules**: Critical guidelines for handling files, project structure, and user interaction
- **System Information**: Details about your environment (OS, shell, working directory)
- **Custom Instructions**: Your global and mode-specific customizations

The system prompt is generated dynamically each time you interact with Roo, adapting to your current mode, available tools, and custom settings.

### Custom System Prompts

Advanced users can create custom system prompts for specific modes by placing a `.roo/system-prompt-<mode_slug>` file in their workspace. When present, Roo uses this file instead of generating the standard system prompt sections, allowing for complete customization of Roo's behavior in that mode.

---

## User Messages

User messages contain your direct inputs to Roo, plus additional contextual information:

- **Your Query**: The text you type in the chat interface
- **Images**: Any images you include in your message (for supported models)
- **Environment Details**: Automatically appended information about your workspace state:
  - Open files/tabs
  - Cursor position
  - Active terminals with output
  - Recently modified files
  - Current time
  - Token/cost information
  - Current mode
  - File listing (on initial connection)

This automatic context enrichment helps Roo understand your workspace without requiring you to explicitly describe it.

---

## Assistant Messages

Assistant messages are the LLM's responses, which may include:

- **Text Responses**: Direct answers to your queries
- **Thinking**: Internal reasoning process (visible when enabled)
- **Tool Calls**: Requests to use specific tools like reading files or executing commands

Note that while assistant messages contain tool calls, the results of those tools are sent back to the LLM in separate tool messages, not as part of the assistant message itself.

---

## Message Flow

Here's how these components work together:

1. **Initial Setup**: Roo generates the system prompt based on your selected mode and configuration
2. **User Input**: You send a message, which is enriched with environment details
3. **LLM Processing**: The LLM receives all previous messages plus your new input
4. **Assistant Response**: The LLM generates a response, potentially using tools
5. **Tool Execution**: If the LLM requests a tool, Roo executes it and provides the result
6. **Conversation History**: All messages are maintained in a structured history for context

---

## Technical Implementation

Internally, Roo's prompt construction is handled by several components:

- **System Prompt Generation**: The `SYSTEM_PROMPT` function in `src/core/prompts/system.ts` assembles the complete system prompt
- **Section Generators**: Specialized functions create each section of the system prompt
- **Message Transformation**: Provider-specific transformers convert Roo's internal message format to the format required by each LLM API
- **Custom Prompt Loading**: The `loadSystemPromptFile` function checks for and processes custom system prompt files

---

## Support Prompts

Alongside the main chat flow, Roo uses specialized templates for specific code actions:

- **Code Action Prompts**: For commands like "Explain", "Fix", "Improve", or "Add to Context"
- **Template-Based**: Generated from templates in `src/shared/support-prompt.ts`
- **Independent Context**: Often operates without the main chat history
- **Task-Specific Format**: Optimized for the specific code task being performed

These support prompts work outside the normal conversation flow to provide focused assistance for specific coding tasks.

---

## Optimizing Your Interactions

Understanding this structure can help you:

- **Write Better Prompts**: Knowing what context Roo already has helps you avoid redundant information
- **Troubleshoot Issues**: Understanding message flow helps identify where problems might occur
- **Create Custom Modes**: With knowledge of the system prompt structure, you can create more effective custom modes
- **Use Custom System Prompts**: Advanced users can create entirely custom system prompts for specialized use cases

This technical foundation powers all of Roo's capabilities, enabling it to understand your requests and effectively utilize available tools to complete tasks.

## **Rate Limits**

---
description: Understand token usage, cost calculation, and optimization strategies for Roo Code. Learn how to manage API costs and set request limits effectively.
keywords:
  - rate limits
  - API costs
  - token usage
  - cost optimization
  - auto-approval limits
  - API management
image: /img/social-share.jpg
---

# Rate Limits and Costs

Understanding and managing API usage is crucial for a smooth and cost-effective experience with Roo Code. This section explains how to track your token usage and costs. Rate limits, which default to 0 (disabled) and typically don't need adjustment, are now configured per profile; see the [API Configuration Profiles](/features/api-configuration-profiles#creating-a-profile) documentation for details on how to set them if needed.

---

## Token Usage

Roo Code interacts with AI models using tokens. Tokens are essentially pieces of words. The number of tokens used in a request and response affects both the processing time and the cost.

*   **Input Tokens:** These are the tokens in your prompt, including the system prompt, your instructions, and any context provided (e.g., file contents).
*   **Output Tokens:** These are the tokens generated by the AI model in its response.

You can see the number of input and output tokens used for each interaction in the chat history.

---

## Cost Calculation

Most AI providers charge based on the number of tokens used. Pricing varies depending on the provider and the specific model.

Roo Code automatically calculates the estimated cost of each API request based on the configured model's pricing. This cost is displayed in the chat history, next to the token usage.

**Note:**

*   The cost calculation is an *estimate*. The actual cost may vary slightly depending on the provider's billing practices.
*   Some providers may offer free tiers or credits. Check your provider's documentation for details.
*   Some providers offer prompt caching which greatly lowers cost.

### Limiting Auto-Approved Requests

To further help manage API costs and prevent unexpected expenses, Roo Code includes a "Max Requests" setting for auto-approved actions. This allows you to define a specific limit on how many consecutive API calls Roo Code can make without requiring your explicit re-approval during a task.

*   **How it works:** If you set a limit (e.g., 5 requests), Roo Code will perform up to 5 auto-approved API calls. Before making the 6th call, it will pause and prompt you to "Reset and Continue," as shown below.
    <img src="/img/v3.18.0/v3.18.0-1.png" alt="Warning message indicating the auto-approved request limit has been reached." width="600" />
    *Notification when the auto-approved request limit is met.*
*   **Configuration:** This limit is configured within the "Auto-approve actions" settings. You can set a specific number or choose "Unlimited." For detailed steps on configuring this and other auto-approval settings, see the [Auto-Approving Actions documentation](/features/auto-approving-actions).
    <img src="/img/v3.18.0/v3.18.0.png" alt="Setting the Max Requests limit for auto-approved actions in Roo Code settings." width="600" />
    *Setting the "Max Requests" for auto-approved actions.*

This feature provides an additional safeguard, particularly for complex or long-running tasks where multiple API calls might be involved.

---

## Tips for Optimizing Token Usage

*   **Be Concise:** Use clear and concise language in your prompts. Avoid unnecessary words or details.
*   **Provide Only Relevant Context:** Use context mentions (`@file.ts`, `@folder/`) selectively. Only include the files that are directly relevant to the task.
*   **Break Down Tasks:** Divide large tasks into smaller, more focused sub-tasks.
*   **Use Custom Instructions:** Provide custom instructions to guide Roo Code's behavior and reduce the need for lengthy explanations in each prompt.
*   **Choose the Right Model:** Some models are more cost-effective than others. Consider using a smaller, faster model for tasks that don't require the full power of a larger model.
*   **Use Modes:** Different modes can access different tools, for example `Architect` can't modify code, which makes it a safe choice when analyzing a complex codebase, without worrying about accidentally allowing expensive operations.
*   **Disable MCP If Not Used:** If you're not using MCP (Model Context Protocol) features, consider [disabling it in the MCP settings](/features/mcp/using-mcp-in-roo#enabling-or-disabling-mcp-server-creation) to significantly reduce the size of the system prompt and save tokens.

By understanding and managing your API usage, you can use Roo Code effectively and efficiently.

## **Roo Code Nightly**

---
description: Learn how to install and use Roo Code Nightly builds to test the latest features and improvements before official releases.
keywords:
  - Roo Code Nightly
  - prerelease builds
  - beta testing
  - latest features
  - nightly builds
image: /img/social-share.jpg
sidebar_label: Roo Code Nightly
---

# Roo Code Nightly

For users who want the latest features and updates, Roo Code offers a prerelease build called **Roo Code Nightly**, which is automatically published whenever changes are merged into our main development branch. This means it always reflects the latest code. It's designed for users who want to access the newest features and contribute to early testing.

---

## Understanding Nightly vs. Other Prereleases

It's important to distinguish Roo Code Nightly from other types of prerelease software (like beta versions or release candidates):

*   **Roo Code Nightly:**
    *   **Frequency:** Automated builds, generated on each merge to the main development branch.
    *   **Content:** Reflects the very latest merged code.
    *   **Stability:** Highly experimental; may contain bugs or incomplete features.
    *   **Testing:** Minimal or no manual testing before release.
    *   **Audience:** Primarily for developers, internal testing, or users comfortable with potentially unstable software who want to see the absolute latest changes and help identify issues early.

*   **General Prerelease (e.g., Beta, Release Candidate):**
    *   **Frequency:** Released periodically (e.g., weekly, monthly) as specific milestones are met.
    *   **Content:** Represents a more curated and planned intermediate version.
    *   **Stability:** More stable than Nightly builds; often undergoes some manual testing or validation.
    *   **Testing:** Typically receives more focused testing.
    *   **Audience:** Intended for a wider group of early adopters or beta testers to gather feedback before a stable release.

In short: **Nightly = latest code, potentially unstable, frequent updates.** Other prereleases are generally less frequent, more validated, and aimed at broader testing.

---

## Installing Roo Code Nightly

To install Roo Code Nightly:

1.  Open VS Code.
2.  Access Extensions: Click the Extensions icon in the Activity Bar or press `Ctrl+Shift+X` (Windows/Linux) or `Cmd+Shift+X` (macOS).
3.  Search for "Roo Code Nightly".
4.  Select "Roo Code Nightly" by Roo Code and click **Install**.
5.  Reload VS Code if prompted.

<img src="/img/installing/installing-5.png" alt="Roo Code Nightly extension in VS Code Marketplace" width="400" />
*Roo Code Nightly in the VS Code Marketplace.*

You can have both the stable version of Roo Code and Roo Code Nightly installed simultaneously.

## **access_mcp_resource**

---
description: Learn how the access_mcp_resource tool retrieves data from Model Context Protocol servers for additional context in Roo Code tasks.
keywords:
  - access_mcp_resource
  - MCP
  - Model Context Protocol
  - MCP resources
  - Roo Code tools
  - context retrieval
  - API integration
image: /img/social-share.jpg
---

# access_mcp_resource

The `access_mcp_resource` tool retrieves data from resources exposed by connected Model Context Protocol (MCP) servers. It allows Roo to access files, API responses, documentation, or system information that provides additional context for tasks.

---

## Parameters

The tool accepts these parameters:

- `server_name` (required): The name of the MCP server providing the resource
- `uri` (required): The URI identifying the specific resource to access

---

## What It Does

This tool connects to MCP servers and fetches data from their exposed resources. Unlike `use_mcp_tool` which executes actions, this tool specifically retrieves information that serves as context for tasks.

---

## When is it used?

- When Roo needs additional context from external systems
- When Roo needs to access domain-specific data from specialized MCP servers
- When Roo needs to retrieve reference documentation hosted by MCP servers
- When Roo needs to integrate real-time data from external APIs via MCP

---

## Key Features

- Retrieves both text and image data from MCP resources
- Requires user approval before executing resource access
- Uses URI-based addressing to precisely identify resources
- Integrates with the Model Context Protocol SDK
- Displays resource content appropriately based on content type
- Supports timeouts for reliable network operations
- Handles server connection states (connected, connecting, disconnected)
- Discovers available resources from connected servers
- Processes structured response data with metadata
- Handles image content special rendering

---

## Limitations

- Depends on external MCP servers being available and connected
- Limited to the resources provided by connected servers
- Cannot access resources from disabled servers
- Network issues can affect reliability and performance
- Resource access subject to configured timeouts
- URI formats are determined by the specific MCP server implementation
- No offline or cached resource access capabilities

---

## How It Works

When the `access_mcp_resource` tool is invoked, it follows this process:

1. **Connection Validation**:
   - Verifies that an MCP hub is available and initialized
   - Confirms the specified server exists in the connection list
   - Checks if the server is disabled (returns an error if it is)

2. **User Approval**:
   - Presents the resource access request to the user for approval
   - Provides server name and resource URI for user verification
   - Proceeds only if the user approves the resource access

3. **Resource Request**:
   - Uses the Model Context Protocol SDK to communicate with servers
   - Makes a `resources/read` request to the server through the MCP hub
   - Applies configured timeouts to prevent hanging on unresponsive servers

4. **Response Processing**:
   - Receives a structured response with metadata and content arrays
   - Processes text content for display to the user
   - Handles image data specially for appropriate display
   - Returns the processed resource data to Roo for use in the current task

---

## Resource Types

MCP servers can provide two main types of resources:

1. **Standard Resources**:
   - Fixed resources with specific URIs
   - Defined name, description, and MIME type
   - Direct access without parameters
   - Typically represent static data or real-time information

2. **Resource Templates**:
   - Parameterized resources with placeholder values in URIs
   - Allow dynamic resource generation based on provided parameters
   - Can represent queries or filtered views of data
   - More flexible but require additional URI formatting

---

## Examples When Used

- When helping with API development, Roo retrieves endpoint specifications from MCP resources to ensure correct implementation.
- When assisting with data visualization, Roo accesses current data samples from connected MCP servers.
- When working in specialized domains, Roo retrieves technical documentation to provide accurate guidance.
- When generating industry-specific code, Roo references compliance requirements from documentation resources.

---

## Usage Examples

Accessing current weather data:
```
<access_mcp_resource>
<server_name>weather-server</server_name>
<uri>weather://san-francisco/current</uri>
</access_mcp_resource>
```

Retrieving API documentation:
```
<access_mcp_resource>
<server_name>api-docs</server_name>
<uri>docs://payment-service/endpoints</uri>
</access_mcp_resource>
```

Accessing domain-specific knowledge:
```
<access_mcp_resource>
<server_name>knowledge-base</server_name>
<uri>kb://medical/terminology/common</uri>
</access_mcp_resource>
```

Fetching system configuration:
```
<access_mcp_resource>
<server_name>infra-monitor</server_name>
<uri>config://production/database</uri>
</access_mcp_resource>
```

## **apply_diff**

---
description: Master the apply_diff tool for making surgical code changes using fuzzy matching and line hints in Roo Code with multi-file support.
keywords:
  - apply_diff
  - file editing
  - code modifications
  - fuzzy matching
  - diff tool
  - Roo Code tools
  - multi-file edits
image: /img/social-share.jpg
---

# apply_diff

The `apply_diff` tool makes precise, surgical changes to files by specifying exactly what content to replace. It uses a sophisticated strategy for finding and applying changes while maintaining proper code formatting and structure.

---

## Parameters

The tool accepts these parameters:

- `path` (required): The path of the file to modify relative to the current working directory.
- `diff` (required): The search/replace block defining the changes using a format specific to the active diff strategy.
- `start_line` (optional): A hint for where the search content begins. _Note: This top-level parameter appears unused by the current main strategy, which relies on `:start_line:` within the diff content._
- `end_line` (optional): A hint for where the search content ends. _Note: This top-level parameter appears unused by the current main strategy._

---

## What It Does

This tool applies targeted changes to existing files using fuzzy matching guided by line number hints to locate and replace content precisely. Unlike simple search and replace, it identifies the exact block for replacement based on the provided content and location hints.

---

## When is it used?

- When Roo needs to make precise changes to existing code without rewriting entire files.
- When refactoring specific sections of code while maintaining surrounding context.
- When fixing bugs in existing code with surgical precision.
- When implementing feature enhancements that modify only certain parts of a file.

---

## Key Features

- Uses fuzzy matching (Levenshtein distance on normalized strings) guided by a `:start_line:` hint, with configurable confidence thresholds (typically 0.8-1.0).
- Provides context around matches using `BUFFER_LINES` (default 40).
- Performs a middle-out search within a configurable context window (`bufferLines`) around the hinted start line.
- Preserves code formatting and indentation passively by replacing exact blocks.
- Shows changes in a diff view for user review and editing before applying.
- Tracks consecutive errors per file (`consecutiveMistakeCountForApplyDiff`) to prevent repeated failures.
- Validates file access against `.rooignore` rules.
- Handles multi-line edits effectively.

---

## Limitations

- Works best with unique, distinctive code sections for reliable identification.
- Performance can vary with very large files or highly repetitive code patterns.
- Fuzzy matching might occasionally select incorrect locations if content is ambiguous.
- Each diff strategy has specific format requirements.
- Complex edits might require careful strategy selection or manual review.

---

## How It Works

When the `apply_diff` tool is invoked, it follows this process:

1.  **Parameter Validation**: Validates required `path` and `diff` parameters.
2.  **RooIgnore Check**: Validates if the target file path is allowed by `.rooignore` rules.
3.  **File Analysis**: Loads the target file content.
4.  **Match Finding**: Uses a fuzzy matching algorithm (Levenshtein on normalized strings) guided by the `:start_line:` hint within a context window (`BUFFER_LINES`), searching middle-out to locate the target content based on the confidence threshold.
5.  **Change Preparation**: Generates the proposed changes by replacing the identified block.
6.  **User Interaction**:
    *   Displays the changes in a diff view.
    *   Allows the user to review and potentially edit the proposed changes.
    *   Waits for user approval or rejection.
7.  **Change Application**: If approved, applies the changes (potentially including user edits) to the file.
8.  **Error Handling**: If errors occur (e.g., match failure, partial application), increments the `consecutiveMistakeCountForApplyDiff` for the file and reports the failure type.
9. **Feedback**: Returns the result, including any user feedback or error details.

---

## Diff Format Requirements

The `<diff>` parameter requires a specific format supporting one or more changes in a single request. Each change block requires a line number hint for the original content.

*   **Requires**: Exact match for the `SEARCH` block content (within the fuzzy threshold), including whitespace and indentation. The `:start_line:` number hint is mandatory within each block. The `:end_line:` hint is optional (but supported by the parser). Markers like `<<<<<<<` within the file's content must be escaped (`\\`) in the SEARCH block.

Example format for the `<diff>` block:

```diff
<<<<<<< SEARCH
:start_line:10
:end_line:12
-------
    // Old calculation logic
    const result = value * 0.9;
    return result;
=======
    // Updated calculation logic with logging
    console.log(`Calculating for value: ${value}`);
    const result = value * 0.95; // Adjusted factor
    return result;
>>>>>>> REPLACE

<<<<<<< SEARCH
:start_line:25
:end_line:25
-------
    const defaultTimeout = 5000;
=======
    const defaultTimeout = 10000; // Increased timeout
>>>>>>> REPLACE
```

---

## Experimental: Multi-File Edits (`MULTI_FILE_APPLY_DIFF`)

An experimental version of `apply_diff` allows for applying changes to multiple files within a single tool call. This feature is activated by the `MULTI_FILE_APPLY_DIFF` experiment flag.

### Key Features of Experimental Mode

- **Multi-File Operations**: Edit multiple files in one request, streamlining complex refactoring tasks.
- **Batch Approval UI**: When multiple files are targeted, a single UI appears for the user to approve or deny all changes at once, or manage permissions for each file individually.
- **New XML Format**: Introduces a new, more structured XML format using `<args>` and `<file>` tags to handle multiple operations.
- **Backward Compatibility**: The experimental tool remains compatible with the legacy single-file `path` and `diff` parameter format.

### How It Works (Experimental)

1.  **Experiment Check**: The tool first checks if the `MULTI_FILE_APPLY_DIFF` experiment is enabled. If not, it falls back to the legacy `apply_diff` implementation.
2.  **Parameter Parsing**: It parses the new `<args>` XML format to identify all target files and their corresponding diff operations. It can also handle the legacy `path`/`diff` parameters.
3.  **Validation and Approval**:
    *   It validates that all target files exist and are accessible (not blocked by `.rooignore`).
    *   If multiple files are being modified, it presents a **batch approval** dialog to the user.
4.  **Diff Application**: For each approved file, it applies the specified diffs using the `MultiFileSearchReplaceDiffStrategy`. This strategy can apply multiple, non-overlapping changes to a single file.
5.  **Results**: It returns a consolidated result message summarizing the outcome for all attempted operations.

### New Diff Format (Experimental)

The experimental mode uses a new XML structure within the `<apply_diff>` tool call.

- **`<args>`**: A container for all file operations.
- **`<file>`**: A block for each file to be modified. Contains `<path>` and one or more `<diff>` tags.
- **`<path>`**: The relative path to the file.
- **`<diff>`**: A block containing the change.
    - **`<content>`**: The `SEARCH/REPLACE` block.
    - **`<start_line>`**: (Optional) A line number hint.

**Example: Modifying two files in one call**

```xml
<apply_diff>
  <args>
    <file>
      <path>src/services/auth.service.ts</path>
      <diff>
        <content>
<<<<<<< SEARCH
:start_line:50
-------
    const token_expiration = '15m';
>>>>>>> REPLACE
        </content>
      </diff>
    </file>
    <file>
      <path>src/config/auth.config.ts</path>
      <diff>
        <content>
<<<<<<< SEARCH
:start_line:12
-------
    rateLimit: 5,
=======
    rateLimit: 10,
>>>>>>> REPLACE
        </content>
      </diff>
    </file>
  </args>
</apply_diff>

```

## **ask_followup_question**

---
description: Enable interactive communication in Roo Code with the ask_followup_question tool for gathering clarification and user preferences.
keywords:
  - ask_followup_question
  - user interaction
  - interactive communication
  - Roo Code tools
  - clarification
  - user feedback
image: /img/social-share.jpg
---

# ask_followup_question

The `ask_followup_question` tool enables interactive communication by asking specific questions to gather additional information needed to complete tasks effectively.

---

## Parameters

The tool accepts these parameters:

- `question` (required): The specific question to ask the user
- `follow_up` (optional): A list of 2-4 suggested answers that help guide user responses, each within `<suggest>` tags

---

## What It Does

This tool creates a conversational interface between Roo and the user, allowing for gathering clarification, additional details, or user preferences when facing ambiguities or decision points. Each question can include suggested responses to streamline the interaction.

---

## When is it used?

- When critical information is missing from the original request
- When Roo needs to choose between multiple valid implementation approaches
- When technical details or preferences are required to proceed
- When Roo encounters ambiguities that need resolution
- When additional context would significantly improve the solution quality

---

## Key Features

- Provides a structured way to gather specific information without breaking workflow
- Includes suggested answers to reduce user typing and guide responses
- Maintains conversation history and context across interactions
- Supports responses containing images and code snippets
- Available in all modes as part of the "always available" tool set
- Enables direct user guidance on implementation decisions
- Formats responses with `<answer>` tags to distinguish them from regular conversation
- Resets consecutive error counter when used successfully

---

## Limitations

- Limited to asking one specific question per tool use
- Presents suggestions as selectable options in the UI
- Cannot force structured responses – users can still respond freely
- Excessive use can slow down task completion and create a fragmented experience
- Suggested answers must be complete, with no placeholders requiring user edits
- No built-in validation for user responses
- Contains no mechanism to enforce specific answer formats

---

## How It Works

When the `ask_followup_question` tool is invoked, it follows this process:

1. **Parameter Validation**: Validates the required `question` parameter and checks for optional suggestions
   - Ensures question text is provided
   - Parses any suggested answers from the `follow_up` parameter using the `fast-xml-parser` library
   - Normalizes suggestions into an array format even if there's only one suggestion

2. **JSON Transformation**: Converts the XML structure into a standardized JSON format for UI display
   ```typescript
   {
     question: "User's question here",
     suggest: [
       { answer: "Suggestion 1" },
       { answer: "Suggestion 2" }
     ]
   }
   ```

3. **UI Integration**:
   - Passes the JSON structure to the UI layer via the `ask("followup", ...)` method
   - Displays selectable suggestion buttons to the user in the interface
   - Creates an interactive experience for selecting or typing a response

4. **Response Collection and Processing**:
   - Captures user text input and any images included in the response
   - Wraps user responses in `<answer>` tags when returning to the assistant
   - Preserves any images included in the user's response
   - Maintains the conversational context by adding the response to the history
   - Resets the consecutive error counter when the tool is used successfully

5. **Error Handling**:
   - Tracks consecutive mistakes using a counter
   - Resets the counter when the tool is used successfully
   - Provides specific error messages:
     - For missing parameters: "Missing required parameter 'question'"
     - For XML parsing: "Failed to parse operations: [error message]"
     - For invalid format: "Invalid operations xml format"
   - Contains safeguards to prevent tool execution when required parameters are missing
   - Increments consecutive mistake count when errors occur

---

## Workflow Sequence

The question-answer cycle follows this sequence:

1. **Information Gap Recognition**: Roo identifies missing information needed to proceed
2. **Specific Question Creation**: Roo formulates a clear, targeted question
3. **Suggestion Development**: Roo creates relevant suggested answers (optional but recommended)
4. **Tool Invocation**: Assistant invokes the tool with question and optional suggestions
5. **UI Presentation**: Question and suggestions are displayed to the user as interactive elements
6. **User Response**: The user selects a suggestion or provides a custom answer
7. **Message Handling**: System handles both partial and complete messages
   - For streaming responses, processes chunks as they arrive
   - For complete messages, processes the entire response at once
   - Maintains state consistency regardless of message chunking
8. **Response Processing**: System wraps the response in `<answer>` tags and preserves images
9. **Context Integration**: Response is added to the conversation history
10. **Task Continuation**: Roo proceeds with the task using the new information

---

## Examples When Used

- When developing a web application, Roo might ask about preferred styling frameworks (Bootstrap, Tailwind, custom CSS)
- When creating an API, Roo might ask about authentication methods (JWT, OAuth, API keys)
- When refactoring code, Roo might ask about prioritizing performance vs. readability
- When setting up a database, Roo might ask about specific schema design preferences
- When creating a custom feature, Roo might ask about specific behavior expectations
- When troubleshooting errors, Roo might ask about specific environment details

---

## Response Format

When a user responds to a question, the response is formatted with `<answer>` tags for clarity and consistency. This format helps clearly separate the user's answer from other conversation elements.

Example response format:
```
<answer>
Use Tailwind CSS for utility-first styling with maximum flexibility
</answer>
```

Key characteristics of the response format:
- Surrounded by `<answer>` tags
- Preserves formatting of the original response
- Maintains images if included in the response
- Returns the complete text, even for long responses

---

## Usage Examples

Asking about implementation preferences:
```
<ask_followup_question>
<question>Which styling approach would you prefer for this web application?</question>
<follow_up>
<suggest>Use Bootstrap for rapid development with consistent components</suggest>
<suggest>Use Tailwind CSS for utility-first styling with maximum flexibility</suggest>
<suggest>Use vanilla CSS with custom styling for complete control and minimal dependencies</suggest>
</follow_up>
</ask_followup_question>
```


Requesting technical clarification:
```
<ask_followup_question>
<question>What database should this application use for storing user data?</question>
<follow_up>
<suggest>MongoDB for flexible schema and document-based storage</suggest>
<suggest>PostgreSQL for relational data with strong consistency guarantees</suggest>
<suggest>Firebase for real-time updates and simplified backend management</suggest>
<suggest>SQLite for lightweight local storage without external dependencies</suggest>
</follow_up>
</ask_followup_question>
```

Addressing ambiguity in requirements:
```
<ask_followup_question>
<question>How should the application handle user authentication?</question>
<follow_up>
<suggest>Implement email/password authentication with account verification</suggest>
<suggest>Use social login providers (Google, GitHub, etc.) for quick signup</suggest>
<suggest>Implement both email/password and social login options</suggest>
</follow_up>
</ask_followup_question>
```

## **attempt_completion**

---
description: Learn how the attempt_completion tool signals task completion, presents results, and enables iterative refinement in Roo Code.
keywords:
  - attempt_completion
  - task completion
  - result presentation
  - Roo Code tools
  - user feedback
  - task summary
image: /img/social-share.jpg
---

# attempt_completion

The `attempt_completion` tool signals that Roo believes a task is complete and presents results to the user. It provides a summary of what was accomplished, optionally includes a command to demonstrate the result, and supports continued refinement through user feedback.

---

## Parameters

The tool accepts these parameters:

- `result` (required): The final result description summarizing what was accomplished
- `command` (optional): A CLI command to execute to demonstrate the result

---

## What It Does

This tool marks the end of a task by presenting a final summary of what was accomplished. It signals to both the user and the system that the current task is complete. When a command is provided, it can also demonstrate the result through a live execution.

---

## When is it used?

- When Roo believes the user's task has been completed
- When summarizing what changes or improvements were made
- When demonstrating the current result with a command execution
- When providing a checkpoint for user feedback on the solution
- When transitioning from one phase of work to potential refinements

---

## Key Features

- Provides a clear signal that Roo believes the task is complete
- Summarizes accomplishments in a concise message
- Optionally demonstrates results through command execution
- Enables user feedback for further refinements
- Displays results in a special UI format distinct from regular messages
- Captures task completion telemetry for system analytics
- Maintains a structured conversation flow by providing checkpoints
- Supports subtask completion within larger workflows
- Ensures users receive a clear summary of what was done
- Available in all modes as part of the "always available" tool group

---

## Limitations

- Should not be used until previous tool uses are confirmed successful (guideline, not enforced)
- Limited to a single command for result demonstration
- Cannot present multiple command options
- Commands require user approval before execution
- Limited to demonstrating results that can be shown via CLI commands
- Cannot be used for partial task completion or progress updates
- Result formatting strips XML closing tags through internal processing

---

## How It Works

When the `attempt_completion` tool is invoked, it follows this process:

1. **Safety Consideration** (guideline, not enforced):
   - The AI is instructed to confirm previous tool uses were successful
   - This is a best practice rather than a programmatically enforced mechanism

2. **Result Presentation**:
   - Displays the completion message to the user in a special "completion_result" UI format
   - Removes XML closing tags from the result text using the `removeClosingTag` function
   - Presents the result differently than regular messages for visual distinction

3. **Command Execution** (if provided):
   - Requests user approval before executing the command
   - Only executes if the user approves
   - Executes the command using the system's command execution functionality
   - Shows the result of the command to the user

4. **Feedback Collection**:
   - Waits for user feedback on the completion result
   - Processes this feedback and returns it to the AI
   - Enables continued refinement based on user input

5. **Task Completion and Continuation**:
   - Signals the task as completed in the system
   - Captures telemetry data for the completed task
   - For subtasks, offers to finish the subtask and resume the parent task
   - Supports continued conversation through the feedback mechanism

6. **Implementation Integration**:
   - Tool results are parsed through the system's parsing mechanism in `parse-assistant-message.ts`
   - The tool is part of the "ALWAYS_AVAILABLE_TOOLS" constant, making it available in all modes

---

## Result Formatting Guidelines

The result message should follow these guidelines:

- Clearly communicate what was accomplished
- Be concise but complete
- Focus on the value delivered to the user
- Avoid unnecessary pleasantries or filler text
- Maintain a professional, straightforward tone
- Present information in a way that's easy to scan and understand
- Acknowledge that the user may provide feedback for further refinements

Note: The system automatically strips XML closing tags from the result text through the `removeClosingTag` function.

---

## Command Selection Guidelines

When including a command, follow these guidelines:

- Choose commands that visually demonstrate the result
- Prefer commands that show the user what was created or modified
- Examples include:
  * `open index.html` to display a created website
  * `npm start` to launch a development server
  * `python app.py` to run a created application
- Avoid commands that merely print text (like `echo` or `cat`)
- Remember that commands require user approval before execution
- Ensure the command is valid for the user's operating system

---

## Feedback and UI Representation

The `attempt_completion` tool has a unique feedback mechanism:

- Results appear in a special "completion_result" UI format distinct from regular messages
- The system waits for user feedback after presenting the result
- Feedback is processed and returned to Roo for further refinements
- This creates an iterative improvement cycle rather than ending the conversation
- The UI includes special elements for providing feedback
- Results serve as clear checkpoints in the conversation flow

When used within subtasks:
- The system offers to finish the subtask and resume the parent task
- If approved, the subtask is completed with a summary
- The parent task resumes with context from the completed subtask
- This enables complex, nested workflows while maintaining context

---

## Examples When Used

- When creating a website, Roo presents the solution and uses `open index.html` to show the result. The user provides feedback, and Roo refines the design.
- When implementing a feature, Roo summarizes the changes and their impact. The user requests additional test coverage, and Roo continues the task.
- When fixing bugs, Roo explains what was fixed and how the solution addresses the root cause. The user confirms the fix works.
- When setting up a development environment, Roo confirms all components are installed. The user requests additional configuration, and Roo continues.
- When creating a data processing script, Roo summarizes the functionality. The user suggests performance improvements, and Roo optimizes the code.

---

## Usage Examples


Completing a task with just a result summary:
```
<attempt_completion>
<result>
I've updated the CSS for the navigation bar to match the design mockup. The changes include:
1. Updated color scheme to use the new brand colors
2. Added responsive behavior for mobile devices
3. Fixed the alignment issues in the dropdown menus
</result>
</attempt_completion>
```

Completing a task with a result and demonstration command:
```
<attempt_completion>
<result>
I've created a simple React application with the following features:
1. A responsive landing page with the requested hero section
2. A working contact form with validation
3. Navigation between all three requested pages
4. Styling that matches the provided design mockup
</result>
<command>npm start</command>
</attempt_completion>
```

Completing a Python script task with execution:
```
<attempt_completion>
<result>
I've created a data processing script that:
1. Reads the CSV files from the data directory
2. Cleans and normalizes the data as requested
3. Generates the summary statistics in JSON format
4. Outputs the processed data to the output directory
</result>
<command>python process_data.py</command>
</attempt_completion>
```

## **browser_action**

---
description: Automate web interactions in Roo Code using the browser_action tool for navigation, clicking, typing, and visual feedback.
keywords:
  - browser_action
  - web automation
  - Puppeteer
  - browser control
  - Roo Code tools
  - web testing
  - screenshots
image: /img/social-share.jpg
---

# browser_action

The `browser_action` tool enables web automation and interaction via a Puppeteer-controlled browser. It allows Roo to launch browsers, navigate to websites, click elements, type text, and scroll pages with visual feedback through screenshots.

---

## Parameters

The tool accepts these parameters:

- `action` (required): The action to perform:
  * `launch`: Start a new browser session at a URL
  * `click`: Click at specific x,y coordinates
  * `type`: Type text via the keyboard
  * `scroll_down`: Scroll down one page height
  * `scroll_up`: Scroll up one page height
  * `close`: End the browser session
- `url` (optional): The URL to navigate to when using the `launch` action
- `coordinate` (optional): The x,y coordinates for the `click` action (e.g., "450,300")
- `text` (optional): The text to type when using the `type` action

---

## What It Does

This tool creates an automated browser session that Roo can control to navigate websites, interact with elements, and perform tasks that require browser automation. Each action provides a screenshot of the current state, enabling visual verification of the process.

---

## When is it used?

- When Roo needs to interact with web applications or websites
- When testing user interfaces or web functionality
- When capturing screenshots of web pages
- When demonstrating web workflows visually

---

## Key Features

- Provides visual feedback with screenshots after each action and captures console logs
- Supports complete workflows from launching to page interaction to closing
- Enables precise interactions via coordinates, keyboard input, and scrolling
- Maintains consistent browser sessions with intelligent page loading detection
- Operates in two modes: local (isolated Puppeteer instance) or remote (connects to existing Chrome)
- Handles errors gracefully with automatic session cleanup and detailed messages
- Optimizes visual output with support for various formats and quality settings
- Tracks interaction state with position indicators and action history

---

## Browser Modes

The tool operates in two distinct modes:

### Local Browser Mode (Default)
- Downloads and manages a local Chromium instance through Puppeteer
- Creates a fresh browser environment with each launch
- No access to existing user profiles, cookies, or extensions
- Consistent, predictable behavior in a sandboxed environment
- Completely closes the browser when the session ends

### Remote Browser Mode
- Connects to an existing Chrome/Chromium instance running with remote debugging enabled
- Can access existing browser state, cookies, and potentially extensions
- Faster startup as it reuses an existing browser process
- Supports connecting to browsers in Docker containers or on remote machines
- Only disconnects (doesn't close) from the browser when session ends
- Requires Chrome to be running with remote debugging port open (typically port 9222)

---

## Limitations

- While the browser is active, only `browser_action` tool can be used
- Browser coordinates are viewport-relative, not page-relative
- Click actions must target visible elements within the viewport
- Browser sessions must be explicitly closed before using other tools
- Browser window has configurable dimensions (default 900x600)
- Cannot directly interact with browser DevTools
- Browser sessions are temporary and not persistent across Roo restarts
- Works only with Chrome/Chromium browsers, not Firefox or Safari
- Local mode has no access to existing cookies; remote mode requires Chrome with debugging enabled

---

## How It Works

When the `browser_action` tool is invoked, it follows this process:

1. **Action Validation and Browser Management**:
   - Validates the required parameters for the requested action
   - For `launch`: Initializes a browser session (either local Puppeteer instance or remote Chrome)
   - For interaction actions: Uses the existing browser session
   - For `close`: Terminates or disconnects from the browser appropriately

2. **Page Interaction and Stability**:
   - Ensures pages are fully loaded using DOM stability detection via `waitTillHTMLStable` algorithm
   - Executes requested actions (navigation, clicking, typing, scrolling) with proper timing
   - Monitors network activity after clicks and waits for navigation when necessary

3. **Visual Feedback**:
   - Captures optimized screenshots using WebP format (with PNG fallback)
   - Records browser console logs for debugging purposes
   - Tracks mouse position and maintains paginated history of actions

4. **Session Management**:
   - Maintains browser state across multiple actions
   - Handles errors and automatically cleans up resources
   - Enforces proper workflow sequence (launch → interactions → close)

---

## Workflow Sequence

Browser interactions must follow this specific sequence:

1. **Session Initialization**: All browser workflows must start with a `launch` action
2. **Interaction Phase**: Multiple `click`, `type`, and scroll actions can be performed
3. **Session Termination**: All browser workflows must end with a `close` action
4. **Tool Switching**: After closing the browser, other tools can be used

---

## Examples When Used

- When creating a web form submission process, Roo launches a browser, navigates to the form, fills out fields with the `type` action, and clicks submit.
- When testing a responsive website, Roo navigates to the site and uses scroll actions to examine different sections.
- When capturing screenshots of a web application, Roo navigates through different pages and takes screenshots at each step.
- When demonstrating an e-commerce checkout flow, Roo simulates the entire process from product selection to payment confirmation.

---

## Usage Examples

Launching a browser and navigating to a website:
```
<browser_action>
<action>launch</action>
<url>https://example.com</url>
</browser_action>
```

Clicking at specific coordinates (e.g., a button):
```
<browser_action>
<action>click</action>
<coordinate>450,300</coordinate>
</browser_action>
```

Typing text into a focused input field:
```
<browser_action>
<action>type</action>
<text>Hello, World!</text>
</browser_action>
```

Scrolling down to see more content:
```
<browser_action>
<action>scroll_down</action>
</browser_action>
```

Closing the browser session:
```
<browser_action>
<action>close</action>
</browser_action>
```

## **codebase_search**

---
description: Perform intelligent semantic searches across your codebase using AI embeddings to find relevant code by meaning, not just keywords.
keywords:
  - codebase_search
  - semantic search
  - AI embeddings
  - code search
  - Roo Code tools
  - vector search
  - Qdrant
image: /img/social-share.jpg
---

# codebase_search

:::info Setup Required
The `codebase_search` tool is part of the [Codebase Indexing](/features/codebase-indexing) feature. It requires additional setup including an embedding provider and vector database.
:::

The `codebase_search` tool performs semantic searches across your entire codebase using AI embeddings. Unlike traditional text-based search, it understands the meaning of your queries and finds relevant code even when exact keywords don't match.

---

## Parameters

The tool accepts these parameters:

- `query` (required): Natural language search query describing what you're looking for
- `path` (optional): Directory path to limit search scope to a specific part of your codebase

---

## What It Does

This tool searches through your indexed codebase using semantic similarity rather than exact text matching. It finds code blocks that are conceptually related to your query, even if they don't contain the exact words you searched for. Results include relevant code snippets with file paths, line numbers, and similarity scores.

---

## When is it used?

- When Roo needs to find code related to specific functionality across your project
- When looking for implementation patterns or similar code structures
- When searching for error handling, authentication, or other conceptual code patterns
- When exploring unfamiliar codebases to understand how features are implemented
- When finding related code that might be affected by changes or refactoring

---

## Key Features

- **Semantic Understanding**: Finds code by meaning rather than exact keyword matches
- **Cross-Project Search**: Searches across your entire indexed codebase, not just open files
- **Contextual Results**: Returns code snippets with file paths and line numbers for easy navigation
- **Similarity Scoring**: Results ranked by relevance with similarity scores (0-1 scale)
- **Scope Filtering**: Optional path parameter to limit searches to specific directories
- **Intelligent Ranking**: Results sorted by semantic relevance to your query
- **UI Integration**: Results displayed with syntax highlighting and navigation links
- **Performance Optimized**: Fast vector-based search with configurable result limits

---

## Requirements

This tool is only available when the Codebase Indexing feature is properly configured:

- **Feature Configured**: Codebase Indexing must be configured in settings
- **Embedding Provider**: OpenAI API key or Ollama configuration required
- **Vector Database**: Qdrant instance running and accessible
- **Index Status**: Codebase must be indexed (status: "Indexed" or "Indexing")

---

## Limitations

- **Requires Configuration**: Depends on external services (embedding provider + Qdrant)
- **Index Dependency**: Only searches through indexed code blocks
- **Result Limits**: Maximum of 50 results per search to maintain performance
- **Similarity Threshold**: Only returns results above similarity threshold (default: 0.4, configurable)
- **File Size Limits**: Limited to files under 1MB that were successfully indexed
- **Language Support**: Effectiveness depends on Tree-sitter language support

---

## How It Works

When the `codebase_search` tool is invoked, it follows this process:

1. **Availability Validation**:
   - Verifies that the CodeIndexManager is available and initialized
   - Confirms codebase indexing is enabled in settings
   - Checks that indexing is properly configured (API keys, Qdrant URL)
   - Validates the current index state allows searching

2. **Query Processing**:
   - Takes your natural language query and generates an embedding vector
   - Uses the same embedding provider configured for indexing (OpenAI or Ollama)
   - Converts the semantic meaning of your query into a mathematical representation

3. **Vector Search Execution**:
   - Searches the Qdrant vector database for similar code embeddings
   - Uses cosine similarity to find the most relevant code blocks
   - Applies the minimum similarity threshold (default: 0.4, configurable) to filter results
   - Limits results to 50 matches for optimal performance

4. **Path Filtering** (if specified):
   - Filters results to only include files within the specified directory path
   - Uses normalized path comparison for accurate filtering
   - Maintains relevance ranking within the filtered scope

5. **Result Processing and Formatting**:
   - Converts absolute file paths to workspace-relative paths
   - Structures results with file paths, line ranges, similarity scores, and code content
   - Formats for both AI consumption and UI display with syntax highlighting

6. **Dual Output Format**:
   - **AI Output**: Structured text format with query, file paths, scores, and code chunks
   - **UI Output**: JSON format with syntax highlighting and navigation capabilities

---

## Search Query Best Practices

### Effective Query Patterns

**Good: Conceptual and specific**
```xml
<codebase_search>
<query>user authentication and password validation</query>
</codebase_search>
```

**Good: Feature-focused**
```xml
<codebase_search>
<query>database connection pool setup</query>
</codebase_search>
```

**Good: Problem-oriented**
```xml
<codebase_search>
<query>error handling for API requests</query>
</codebase_search>
```

**Less effective: Too generic**
```xml
<codebase_search>
<query>function</query>
</codebase_search>
```

### Query Types That Work Well

- **Functional Descriptions**: "file upload processing", "email validation logic"
- **Technical Patterns**: "singleton pattern implementation", "factory method usage"
- **Domain Concepts**: "user profile management", "payment processing workflow"
- **Architecture Components**: "middleware configuration", "database migration scripts"

---

## Directory Scoping

Use the optional `path` parameter to focus searches on specific parts of your codebase:

**Search within API modules:**
```xml
<codebase_search>
<query>endpoint validation middleware</query>
<path>src/api</path>
</codebase_search>
```

**Search in test files:**
```xml
<codebase_search>
<query>mock data setup patterns</query>
<path>tests</path>
</codebase_search>
```

**Search specific feature directories:**
```xml
<codebase_search>
<query>component state management</query>
<path>src/components/auth</path>
</codebase_search>
```

---

## Result Interpretation

### Similarity Scores

- **0.8-1.0**: Highly relevant matches, likely exactly what you're looking for
- **0.6-0.8**: Good matches with strong conceptual similarity
- **0.4-0.6**: Potentially relevant but may require review
- **Below 0.4**: Filtered out as too dissimilar

### Result Structure

Each search result includes:
- **File Path**: Workspace-relative path to the file containing the match
- **Score**: Similarity score indicating relevance (0.4-1.0)
- **Line Range**: Start and end line numbers for the code block
- **Code Chunk**: The actual code content that matched your query

---

## Examples When Used

- When implementing a new feature, Roo searches for "authentication middleware" to understand existing patterns before writing new code.
- When debugging an issue, Roo searches for "error handling in API calls" to find related error patterns across the codebase.
- When refactoring code, Roo searches for "database transaction patterns" to ensure consistency across all database operations.
- When onboarding to a new codebase, Roo searches for "configuration loading" to understand how the application bootstraps.

---

## Usage Examples

Searching for authentication-related code across the entire project:
```xml
<codebase_search>
<query>user login and authentication logic</query>
</codebase_search>
```

Finding database-related code in a specific directory:
```xml
<codebase_search>
<query>database connection and query execution</query>
<path>src/data</path>
</codebase_search>
```

Looking for error handling patterns in API code:
```xml
<codebase_search>
<query>HTTP error responses and exception handling</query>
<path>src/api</path>
</codebase_search>
```

Searching for testing utilities and mock setups:
```xml
<codebase_search>
<query>test setup and mock data creation</query>
<path>tests</path>
</codebase_search>
```

Finding configuration and environment setup code:
```xml
<codebase_search>
<query>environment variables and application configuration</query>
</codebase_search>
```

## **execute_command**

---
description: Execute terminal commands in Roo Code for system operations, dependency installation, builds, and development workflows.
keywords:
  - execute_command
  - CLI commands
  - terminal
  - system operations
  - Roo Code tools
  - command execution
  - shell integration
image: /img/social-share.jpg
---

# execute_command

The `execute_command` tool runs CLI commands on the user's system. It allows Roo to perform system operations, install dependencies, build projects, start servers, and execute other terminal-based tasks needed to accomplish user objectives.

---

## Parameters

The tool accepts these parameters:

- `command` (required): The CLI command to execute. Must be valid for the user's operating system.
- `cwd` (optional): The working directory to execute the command in. If not provided, the current working directory is used.

---

## What It Does

This tool executes terminal commands directly on the user's system, enabling a wide range of operations from file manipulations to running development servers. Commands run in managed terminal instances with real-time output capture, integrated with VS Code's terminal system for optimal performance and security.

---

## When is it used?

- When installing project dependencies (npm install, pip install, etc.)
- When building or compiling code (make, npm run build, etc.)
- When starting development servers or running applications
- When initializing new projects (git init, npm init, etc.)
- When performing file operations beyond what other tools provide
- When running tests or linting operations
- When needing to execute specialized commands for specific technologies

---

## Key Features

- Integrates with VS Code shell API for reliable terminal execution
- Reuses terminal instances when possible through a registry system
- Captures command output line by line with real-time feedback
- Supports long-running commands that continue in the background
- Allows specification of custom working directories
- Maintains terminal history and state across command executions
- Handles complex command chains appropriate for the user's shell
- Provides detailed command completion status and exit code interpretation
- Supports interactive terminal applications with user feedback loop
- Shows terminals during execution for transparency
- Validates commands for security using shell-quote parsing
- Blocks potentially dangerous subshell execution patterns
- Integrates with RooIgnore system for file access control
- Handles terminal escape sequences for clean output

---

## Limitations

- Command access may be restricted by RooIgnore rules and security validations
- Commands with elevated permission requirements may need user configuration
- Behavior may vary across operating systems for certain commands
- Very long-running commands may require specific handling
- File paths should be properly escaped according to the OS shell rules
- Not all terminal features may work with remote development scenarios

---

## How It Works

When the `execute_command` tool is invoked, it follows this process:

1. **Command Validation and Security Checks**:
   - Parses the command using shell-quote to identify components
   - Validates against security restrictions (subshell usage, restricted files)
   - Checks against RooIgnore rules for file access permissions
   - Ensures the command meets system security requirements

2. **Terminal Management**:
   - Gets or creates a terminal through TerminalRegistry
   - Sets up the working directory context
   - Prepares event listeners for output capture
   - Shows the terminal for user visibility

3. **Command Execution and Monitoring**:
   - Executes via VS Code's shellIntegration API
   - Captures output with escape sequence processing
   - Throttles output handling (100ms intervals)
   - Monitors for command completion or errors
   - Detects "hot" processes like compilers for special handling

4. **Result Processing**:
   - Strips ANSI/VS Code escape sequences for clean output
   - Interprets exit codes with detailed signal information
   - Updates working directory tracking if changed by command
   - Provides command status with appropriate context

---

## Terminal Implementation Details

The tool uses a sophisticated terminal management system:

1. **First Priority: Terminal Reuse**
   - The TerminalRegistry tries to reuse existing terminals when possible
   - This reduces proliferation of terminal instances and improves performance
   - Terminal state (working directory, history) is preserved across commands

2. **Second Priority: Security Validation**
   - Commands are parsed using shell-quote for component analysis
   - Dangerous patterns like `$(...)` and backticks are blocked
   - Commands are checked against RooIgnore rules for file access control
   - A prefix-based allowlist system validates command patterns

3. **Performance Optimizations**
   - Output is processed in 100ms throttled intervals to prevent UI overload
   - Zero-copy buffer management uses index-based tracking for efficiency
   - Special handling for compilation and "hot" processes
   - Platform-specific optimizations for Windows PowerShell

4. **Error and Signal Handling**
   - Exit codes are mapped to detailed signal information (SIGTERM, SIGKILL, etc.)
   - Core dump detection for critical failures
   - Working directory changes are tracked and handled automatically
   - Clean recovery from terminal disconnection scenarios

---

## Examples When Used

- When setting up a new project, Roo runs initialization commands like `npm init -y` followed by installing dependencies.
- When building a web application, Roo executes build commands like `npm run build` to compile assets.
- When deploying code, Roo runs git commands to commit and push changes to a repository.
- When troubleshooting, Roo executes diagnostic commands to gather system information.
- When starting a development server, Roo launches the appropriate server command (e.g., `npm start`).
- When running tests, Roo executes the test runner command for the project's testing framework.

---

## Usage Examples

Running a simple command in the current directory:
```
<execute_command>
<command>npm run dev</command>
</execute_command>
```

Installing dependencies for a project:
```
<execute_command>
<command>npm install express mongodb mongoose dotenv</command>
</execute_command>
```

Running multiple commands in sequence:
```
<execute_command>
<command>mkdir -p src/components && touch src/components/App.js</command>
</execute_command>
```

Executing a command in a specific directory:
```
<execute_command>
<command>git status</command>
<cwd>./my-project</cwd>
</execute_command>
```

Building and then starting a project:
```
<execute_command>
<command>npm run build && npm start</command>
</execute_command>
```

## **insert_content**

---
description: Insert new content at specific line numbers in existing files without modifying original content using Roo Code's insert_content tool.
keywords:
  - insert_content
  - file insertion
  - add content
  - Roo Code tools
  - code insertion
  - file editing
image: /img/social-share.jpg
---

# insert_content

The `insert_content` tool adds new lines of content into an existing file without modifying the original content. It's ideal for inserting code blocks, configuration entries, or log lines at specific locations.

---

## Parameters

The tool accepts these parameters:

- `path` (required): The relative path (from the workspace root) of the file to insert content into.
- `line` (required): The 1-based line number *before* which the content should be inserted. Use `0` to append the content to the end of the file.
- `content` (required): The text content to insert.

---

## What It Does

This tool reads the target file, identifies the specified insertion point based on the `line` parameter, and inserts the provided `content` at that location. If `line` is `0`, the content is added to the end. Changes are presented in a diff view for user approval before being saved.

---

## When is it used?

- When adding new import statements at the beginning of a file.
- When inserting new functions or methods into existing code.
- When adding configuration blocks to settings files.
- When appending log entries or data records.
- When adding any multi-line text block without altering existing lines.

---

## Key Features

- **Targeted Insertion**: Adds content precisely at the specified line number or appends to the end.
- **Preserves Existing Content**: Does not modify or delete original file lines.
- **Interactive Approval**: Shows proposed insertions in a diff view, requiring explicit user approval.
- **User Edit Support**: Allows editing the proposed content directly within the diff view before final approval.
- **Handles Line Numbers**: Correctly interprets the `line` parameter (1-based or 0 for append).
- **Context Tracking**: Records the file edit operation for context management.
- **Error Handling**: Checks for missing parameters, invalid line numbers, and file access issues.

---

## Limitations

- **Insert Only**: Cannot replace or delete existing content. Use `apply_diff` or `search_and_replace` for modifications.
- **Requires Existing File**: The target file specified by `path` must exist.
- **Review Overhead**: The mandatory diff view approval adds an interactive step.

---

## How It Works

When the `insert_content` tool is invoked, it follows this process:

1.  **Parameter Validation**: Checks for required `path`, `line`, and `content`. Validates `line` is a non-negative integer.
2.  **File Reading**: Reads the content of the target file specified by `path`.
3.  **Insertion Point Calculation**: Converts the 1-based `line` parameter to a 0-based index for internal processing (`-1` for appending).
4.  **Content Insertion**: Uses an internal utility (`insertGroups`) to merge the original file lines with the new `content` at the calculated index.
5.  **Diff View Interaction**:
    *   Opens the file in the diff view (`cline.diffViewProvider.open`).
    *   Updates the diff view with the proposed content (`cline.diffViewProvider.update`).
6.  **User Approval**: Presents the change via `askApproval`. Reverts if rejected.
7.  **Saving Changes**: If approved, saves the changes using `cline.diffViewProvider.saveChanges`.
8.  **File Context Tracking**: Tracks the edit using `cline.getFileContextTracker().trackFileContext`.
9.  **Handling User Edits**: If the user edited the content in the diff view, reports the final merged content back.
10. **Result Reporting**: Reports success (including user edits) or failure back to the AI model.

---

## Usage Examples

Inserting import statements at the beginning of a file (`line: 1`):

```xml
<insert_content>
<path>src/utils.ts</path>
<line>1</line>
<content>
// Add imports at start of file
import { sum } from './math';
import { parse } from 'date-fns';
</content>
</insert_content>
```

Appending content to the end of a file (`line: 0`):

```xml
<insert_content>
<path>config/routes.yaml</path>
<line>0</line>
<content>
- path: /new-feature
  component: NewFeatureComponent
  auth_required: true
</content>
</insert_content>
```

Inserting a function before line 50:

```xml
<insert_content>
<path>src/services/api.js</path>
<line>50</line>
<content>
async function fetchUserData(userId) {
  const response = await fetch(`/api/users/${userId}`);
  if (!response.ok) {
    throw new Error('Failed to fetch user data');
  }
  return response.json();
}
</content>
</insert_content>
```

## **list_code_definition_names**

---
description: Get a structural overview of your codebase by listing classes, functions, and interfaces with the list_code_definition_names tool.
keywords:
  - list_code_definition_names
  - code structure
  - code definitions
  - Roo Code tools
  - codebase overview
  - Tree-sitter
image: /img/social-share.jpg
---

# list_code_definition_names

The `list_code_definition_names` tool provides a structural overview of your codebase by listing code definitions from source files at the top level of a specified directory. It helps Roo understand code architecture by displaying line numbers and definition snippets.

---

## Parameters

The tool accepts these parameters:

- `path` (required): The path of the directory to list top level source code definitions for, relative to the current working directory

---

## What It Does

This tool scans source code files at the top level of a specified directory and extracts code definitions like classes, functions, and interfaces. It displays the line numbers and actual code for each definition, providing a quick way to map the important components of your codebase.

---

## When is it used?

- When Roo needs to understand your codebase architecture quickly
- When Roo needs to locate important code constructs across multiple files
- When planning refactoring or extensions to existing code
- Before diving into implementation details with other tools
- When identifying relationships between different parts of your codebase

---

## Key Features

- Extracts classes, functions, methods, interfaces, and other definitions from source files
- Displays line numbers and actual source code for each definition
- Supports multiple programming languages including JavaScript, TypeScript, Python, Rust, Go, C++, C, C#, Ruby, Java, PHP, Swift, and Kotlin
- Processes only files at the top level of the specified directory (not subdirectories)
- Limits processing to a maximum of 50 files for performance
- Focuses on top-level definitions to avoid overwhelming detail
- Helps identify code organization patterns across the project
- Creates a mental map of your codebase's architecture
- Works in conjunction with other tools like `read_file` for deeper analysis

---

## Limitations

- Only identifies top-level definitions, not nested ones
- Only processes files at the top level of the specified directory, not subdirectories
- Limited to processing a maximum of 50 files per request
- Dependent on language-specific parsers, with varying detection quality
- May not recognize all definitions in languages with complex syntax
- Not a substitute for reading code to understand implementation details
- Cannot detect runtime patterns or dynamic code relationships
- Does not provide information about how definitions are used
- May have reduced accuracy with highly dynamic or metaprogrammed code
- Limited to specific languages supported by the implemented Tree-sitter parsers

---

## How It Works

When the `list_code_definition_names` tool is invoked, it follows this process:

1. **Parameter Validation**: Validates the required `path` parameter
2. **Path Resolution**: Resolves the relative path to an absolute path
3. **Directory Scanning**: Scans only the top level of the specified directory for source code files (not recursive)
4. **File Filtering**: Limits processing to a maximum of 50 files
5. **Language Detection**: Identifies file types based on extensions (.js, .jsx, .ts, .tsx, .py, .rs, .go, .cpp, .hpp, .c, .h, .cs, .rb, .java, .php, .swift, .kt, .kts)
6. **Code Parsing**: Uses Tree-sitter to parse code and extract definitions through these steps:
   - Parsing file content into an Abstract Syntax Tree (AST)
   - Creating a query using a language-specific query string
   - Sorting the captures by their position in the file
7. **Result Formatting**: Outputs definitions with line numbers and actual source code

---

## Output Format

The output shows file paths followed by line numbers and the actual source code of each definition. For example:

```
src/utils.js:
0--0 | export class HttpClient {
5--5 | formatDate() {
10--10 | function parseConfig(data) {

src/models/User.js:
0--0 | interface UserProfile {
10--10 | export class User {
20--20 | function createUser(data) {
```

Each line displays:
- The start and end line numbers of the definition
- The pipe symbol (|) as a separator
- The actual source code of the definition

This output format helps you quickly see both where definitions are located in the file and their implementation details.

---

## Examples When Used

- When starting a new task, Roo first lists key code definitions to understand the overall structure of your project.
- When planning refactoring work, Roo uses this tool to identify classes and functions that might be affected.
- When exploring unfamiliar codebases, Roo maps the important code constructs before diving into implementation details.
- When adding new features, Roo identifies existing patterns and relevant code definitions to maintain consistency.
- When troubleshooting bugs, Roo maps the codebase structure to locate potential sources of the issue.
- When planning architecture changes, Roo identifies all affected components across files.

---

## Usage Examples

Listing code definitions in the current directory:
```
<list_code_definition_names>
<path>.</path>
</list_code_definition_names>
```

Examining a specific module's structure:
```
<list_code_definition_names>
<path>src/components</path>
</list_code_definition_names>
```

Exploring a utility library:
```
<list_code_definition_names>
<path>lib/utils</path>
</list_code_definition_names>
```

## **list_files**

---
description: Learn how the list_files tool helps Roo Code explore project structures, list directories, and navigate codebases with recursive and filtered listing capabilities.
keywords:
  - list_files
  - Roo Code tools
  - directory listing
  - file exploration
  - project structure
  - recursive listing
  - codebase navigation
  - VS Code AI
image: /img/social-share.jpg
---

# list_files

The `list_files` tool displays the files and directories within a specified location. It helps Roo understand your project structure and navigate your codebase effectively.

---

## Parameters

The tool accepts these parameters:

- `path` (required): The path of the directory to list contents for, relative to the current working directory
- `recursive` (optional): Whether to list files recursively. Use `true` for recursive listing, `false` or omit for top-level only.

---

## What It Does

This tool lists all files and directories in a specified location, providing a clear overview of your project structure. It can either show just the top-level contents or recursively explore subdirectories.

---

## When is it used?

- When Roo needs to understand your project structure
- When Roo explores what files are available before reading specific ones
- When Roo maps a codebase to better understand its organization
- Before using more targeted tools like `read_file` or `search_files`
- When Roo needs to check for specific file types (like configuration files) across a project

---

## Key Features

- Lists both files and directories with directories clearly marked
- Offers both recursive and non-recursive listing modes
- Intelligently ignores common large directories like `node_modules` and `.git` in recursive mode
- Respects `.gitignore` rules when in recursive mode
- Marks files ignored by `.rooignore` with a lock symbol (🔒) when `showRooIgnoredFiles` is enabled
- Optimizes file listing performance by leveraging the `ripgrep` tool.
- Sorts results to show directories before their contents, maintaining a logical hierarchy
- Presents results in a clean, organized format
- Automatically creates a mental map of your project structure

---

## Limitations

- File listing is capped at about 200 files by default to prevent performance issues
- The underlying `ripgrep` file listing process has a 10-second timeout; if exceeded, partial results may be returned.
- When the file limit is hit, it adds a note suggesting to use `list_files` on specific subdirectories
- Not designed for confirming the existence of files you've just created
- May have reduced performance in very large directory structures
- Cannot list files in root or home directories for security reasons

---

## How It Works

When the `list_files` tool is invoked, it follows this process:

1. **Parameter Validation**: Validates the required `path` parameter and optional `recursive` parameter
2. **Path Resolution**: Resolves the relative path to an absolute path
3. **Security Checks**: Prevents listing files in sensitive locations like root or home directories
4. **Directory/File Scanning**:
   - Uses the `ripgrep` tool to efficiently list files, applying a 10-second timeout.
   - Uses Node.js `fs` module to list directories.
   - Applies different filtering logic for recursive vs. non-recursive modes.
5. **Result Filtering**:
   - In recursive mode, skips common large directories like `node_modules`, `.git`, etc.
   - Respects `.gitignore` rules when in recursive mode
   - Handles `.rooignore` patterns, either hiding files or marking them with a lock symbol
6. **Formatting**:
   - Marks directories with a trailing slash (`/`)
   - Sorts results to show directories before their contents for logical hierarchy
   - Marks ignored files with a lock symbol (🔒) when `showRooIgnoredFiles` is enabled
   - Caps results at 200 files by default with a note about using subdirectories
   - Organizes results for readability

---

## File Listing Format

The file listing results include:

- Each file path is displayed on its own line
- Directories are marked with a trailing slash (`/`)
- Files ignored by `.rooignore` are marked with a lock symbol (🔒) when `showRooIgnoredFiles` is enabled
- Results are sorted logically with directories appearing before their contents
- When the file limit is reached, a message appears suggesting to use `list_files` on specific subdirectories

Example output format:
```
src/
src/components/
src/components/Button.tsx
src/components/Header.tsx
src/utils/
src/utils/helpers.ts
src/index.ts
...
File listing truncated (showing 200 of 543 files). Use list_files on specific subdirectories for more details.
```

When `.rooignore` files are used and `showRooIgnoredFiles` is enabled:
```
src/
src/components/
src/components/Button.tsx
src/components/Header.tsx
🔒 src/secrets.json
src/utils/
src/utils/helpers.ts
src/index.ts
```

---

## Examples When Used

- When starting a new task, Roo may list the project files to understand its structure before diving into specific code.
- When asked to find specific types of files (like all JavaScript files), Roo first lists directories to know where to look.
- When providing recommendations for code organization, Roo examines the current project structure first.
- When setting up a new feature, Roo lists related directories to understand the project conventions.

---

## Usage Examples

Listing top-level files in the current directory:
```
<list_files>
<path>.</path>
</list_files>
```

Recursively listing all files in a source directory:
```
<list_files>
<path>src</path>
<recursive>true</recursive>
</list_files>
```

Examining a specific project subdirectory:
```
<list_files>
<path>src/components</path>
<recursive>false</recursive>
</list_files>
```

## **new_task**

---
description: Discover how the new_task tool enables complex workflow management by creating subtasks with different modes, maintaining parent-child relationships for organized development.
keywords:
  - new_task
  - Roo Code tools
  - subtasks
  - workflow management
  - task hierarchy
  - mode switching
  - complex projects
  - task organization
  - VS Code AI
image: /img/social-share.jpg
---

# new_task

The `new_task` tool creates subtasks with specialized modes while maintaining a parent-child relationship. It breaks down complex projects into manageable pieces, each operating in the mode best suited for specific work.

---

## Parameters

The tool accepts these parameters:

- `mode` (required): The slug of the mode to start the new task in (e.g., "code", "ask", "architect")
- `message` (required): The initial user message or instructions for this new task
- `todos` (optional): Initial todo list in markdown checklist format

---

## What It Does

This tool creates a new task instance with a specified starting mode and initial message. It allows complex workflows to be divided into subtasks with their own conversation history. Parent tasks are paused during subtask execution and resumed when the subtask completes, with results transferred back to the parent.

---

## When is it used?

- When breaking down complex projects into separate, focused subtasks
- When different aspects of a task require different specialized modes
- When different phases of work benefit from context separation
- When organizing multi-phase development workflows

---

## Key Features

- Creates subtasks with their own conversation history and specialized mode
- Pauses parent tasks for later resumption
- Maintains hierarchical task relationships for navigation
- Transfers results back to parent tasks upon completion
- Supports workflow segregation for complex projects
- Allows different parts of a project to use modes optimized for specific work
- Requires explicit user approval for task creation
- Provides clear task transition in the UI

---

## Limitations

- Cannot create tasks with modes that don't exist
- Requires user approval before creating each new task
- Task interface may become complex with deeply nested subtasks
- Subtasks inherit certain workspace and extension configurations from parents
- May require re-establishing context when switching between deeply nested tasks
- Task completion needs explicit signaling to properly return to parent tasks

---

## How It Works

When the `new_task` tool is invoked, it follows this process:

1. **Parameter Validation**:
   - Validates the required `mode` and `message` parameters
   - Verifies that the requested mode exists in the system

2. **Task Stack Management**:
   - Maintains a task stack that tracks all active and paused tasks
   - Preserves the current mode for later resumption
   - Sets the parent task to paused state

3. **Task Context Management**:
   - Creates a new task context with the provided message
   - Assigns unique taskId and instanceId identifiers for state management
   - Captures telemetry data on tool usage and task lifecycles

4. **Mode Switching and Integration**:
   - Switches to the specified mode with appropriate role and capabilities
   - Initializes the new task with the provided message
   - Integrates with VS Code's command palette and code actions

5. **Task Completion and Result Transfer**:
   - When subtask completes, result is passed back to parent task via `finishSubTask()`
   - Parent task resumes in its original mode
   - Task history and token usage metrics are updated
   - The `taskCompleted` event is emitted with performance data

---

## Configuration

Streamline hierarchical task planning with the optional todo list parameter for subtasks:

- **Pass Todo Lists**: Include predefined todo lists when creating subtasks
- **Maintain Context**: Pass along context to the subtask in the form of a todo list
- **Optional Enforcement**: The "New Task Require Todos" setting in VS Code can enforce todo lists for all new subtasks if desired

<img src="/img/v3.25.21/v3.25.21.png" alt="Subtask todo lists configuration in VS Code settings" width="600" />

This feature works out of the box, and you can optionally configure VS Code settings to require todos for all new tasks.

---

## Examples When Used

- When a front-end developer needs to architect a new feature, implement the code, and document it, they can create separate tasks for each phase with results flowing from one phase to the next.
- When debugging an issue before implementing a fix, the debugging task can document findings that are passed to the implementation task.
- When developing a full-stack application, database schema designs from an architect-mode task inform implementation details in a subsequent code-mode task.
- When documenting a system after implementation, the documentation task can reference the completed implementation while using documentation-specific features.

---

## Usage Examples

Creating a new task in code mode:
```
<new_task>
<mode>code</mode>
<message>Implement a user authentication service with login, registration, and password reset functionality.</message>
</new_task>
```

Creating a documentation task after completing implementation:
```
<new_task>
<mode>docs</mode>
<message>Create comprehensive API documentation for the authentication service we just built.</message>
</new_task>
```

Breaking down a complex feature into architectural planning and implementation:
```
<new_task>
<mode>architect</mode>
<message>Design the database schema and system architecture for our new e-commerce platform.</message>
</new_task>
```

Creating a task with an initial todo list:
```
<new_task>
<mode>code</mode>
<message>Build a REST API for user management</message>
<todos>
[ ] Set up Express server
[ ] Create user model
[ ] Implement CRUD endpoints
[ ] Add authentication middleware
[ ] Write API tests
</todos>
</new_task>
```

## **read_file**

---
description: Explore the read_file tool's capabilities for examining file contents, supporting line ranges, PDF/DOCX extraction, image reading, and experimental multi-file concurrent reading.
keywords:
  - read_file
  - Roo Code tools
  - file reading
  - concurrent reads
  - line numbers
  - PDF extraction
  - DOCX support
  - image support
  - OCR workflows
  - code analysis
  - VS Code AI
image: /img/social-share.jpg
---

# read_file

The `read_file` tool examines the contents of files in a project. It allows Roo to understand code, configuration files, documentation, and now images to provide better assistance.

:::info Multi-File Support
The `read_file` tool can read multiple files simultaneously when the `maxConcurrentFileReads` setting is greater than 1. This significantly improves efficiency for tasks requiring analysis of multiple related files.

**Note:** When reading files (even single files), the LLM will see a message encouraging multi-file reads: "Reading multiple files at once is more efficient for the LLM. If other files are relevant to your current task, please read them simultaneously."
:::

---

## Parameters

The tool accepts parameters in two formats depending on your configuration:

### Standard Format (Single File)

- `path` (required): The path of the file to read relative to the current working directory
- `start_line` (optional): The starting line number to read from (1-based indexing)
- `end_line` (optional): The ending line number to read to (1-based, inclusive)

:::note Legacy Format
While the single-file parameters (`path`, `start_line`, `end_line`) are still supported for backward compatibility, we recommend using the newer `args` format for consistency and future compatibility.
:::

### Enhanced Format (Multi-File)

When `maxConcurrentFileReads` is set to a value greater than 1 (found in Settings > Context > "Concurrent file reads limit"), the tool accepts an `args` parameter containing multiple file entries:

- `args` (required): Container for multiple file specifications
  - `file` (required): Individual file specification
    - `path` (required): The path of the file to read
    - `line_range` (optional): Line range specification (e.g., "1-50" or "100-150"). Multiple `line_range` elements can be specified per file.

---

## What It Does

This tool reads the content of a specified file and returns it with line numbers for easy reference. It can read entire files or specific sections, extract text from PDFs and Word documents, and display images in various formats.

---

## When is it used?

- When Roo needs to understand existing code structure
- When Roo needs to analyze configuration files
- When Roo needs to extract information from text files
- When Roo needs to see code before suggesting changes
- When specific line numbers need to be referenced in discussions

---

## Key Features

- Displays file content with line numbers for easy reference
- Can read specific portions of files by specifying line ranges
- Extracts readable text from PDF and DOCX files
- **Image support**: Displays images in multiple formats (PNG, JPG, JPEG, GIF, WebP, SVG, BMP, ICO, TIFF)
- Automatically truncates large text files when no line range is specified, showing the beginning of the file
- Provides method summaries with line ranges for truncated large code files
- Efficiently streams only requested line ranges for better performance
- Makes it easy to discuss specific parts of code with line numbering
- **Multi-file support**: Read multiple files simultaneously with batch approval (when `maxConcurrentFileReads` > 1)

---

## Multi-File Capabilities

When the `maxConcurrentFileReads` setting is greater than 1, the `read_file` tool gains enhanced capabilities:

### Configuration
- **Location**: Settings > Context > "Concurrent file reads limit"
- **Description**: "Maximum number of files the 'read_file' tool can process concurrently. Higher values may speed up reading multiple small files but increase memory usage."
- **Range**: 1-100 (slider control)
- **Default**: 5

### Batch Processing
- Read up to 100 files in a single request (configurable, default 5)
- Parallel processing for improved performance
- Batch approval interface for user consent

### Enhanced User Experience
- Single approval dialog for multiple files
- Individual file override options
- Clear visibility into which files will be accessed
- Graceful handling of mixed success/failure scenarios

### Improved Efficiency
- Reduces interruptions from multiple approval dialogs
- Faster processing through parallel file reading
- Smart batching of related files
- Configurable concurrency limits to match system capabilities

---

## Limitations

- May not handle extremely large files efficiently without using line range parameters
- For binary files (except PDF, DOCX, and supported image formats), may return content that isn't human-readable
- **Multi-file mode**: Requires `maxConcurrentFileReads` > 1 in settings
- **Image files**: Returns base64-encoded data URLs which may be large for high-resolution images
  - Default max single image size: 20MB
  - Default max total image size: 20MB

---

## How It Works

When the `read_file` tool is invoked, it follows this process:

1. **Parameter Validation**: Validates the required `path` parameter and optional parameters
2. **Path Resolution**: Resolves the relative path to an absolute path
3. **Reading Strategy Selection**:
   - The tool uses a strict priority hierarchy (explained in detail below)
   - It chooses between range reading, auto-truncation, or full file reading
4. **Content Processing**:
   - Adds line numbers to the content (e.g., "1 | const x = 13") where `1 |` is the line number.
   - For truncated files, adds truncation notice and method definitions
   - For special formats (PDF, DOCX), extracts readable text
   - For image formats, returns base64-encoded data URLs with MIME type

---

## Reading Strategy Priority

The tool uses a clear decision hierarchy to determine how to read a file:

1. **First Priority: Explicit Line Range**
   - If either `start_line` or `end_line` is provided, the tool always performs a range read
   - The implementation efficiently streams only the requested lines, making it suitable for processing large files
   - This takes precedence over all other options

2. **Second Priority: Automatic Truncation for Large Text Files**
   - This applies only when **all** of the following conditions are met:
     - Neither `start_line` nor `end_line` is specified.
     - The file is identified as a text-based file (not binary like PDF/DOCX).
     - The file's total line count exceeds the `maxReadFileLine` setting (default: 500 lines).
   - When automatic truncation occurs:
     - The tool reads only the *first* `maxReadFileLine` lines.
     - It appends a notice indicating truncation (e.g., `[Showing only 500 of 1200 total lines...]`).
     - For code files, it may also append a summary of source code definitions found within the truncated portion.
   - **Special Case - Definitions Only Mode**: When `maxReadFileLine` is set to `0`, the tool returns only source code definitions without any file content.

3. **Default Behavior: Read Entire File**
    - If neither an explicit range is given nor automatic truncation applies (e.g., the file is within the line limit, or it's a supported binary type), the tool reads the entire content.
    - For supported formats like PDF and DOCX, it attempts to extract the full text content.
    - For image formats, it returns a base64-encoded data URL that can be displayed in the chat interface.

---

## Examples When Used

- When asked to explain or improve code, Roo first reads the relevant files to understand the current implementation.
- When troubleshooting configuration issues, Roo reads config files to identify potential problems.
- When working with documentation, Roo reads existing docs to understand the current content before suggesting improvements.

---

## Usage Examples

Here are several scenarios demonstrating how the `read_file` tool is used and the typical output you might receive.

### Reading an Entire File

To read the complete content of a file:

**Input:**
```xml
<read_file>
<path>src/app.js</path>
</read_file>
```

**Simulated Output (for a small file like `example_small.txt`):**
```
1 | This is the first line.
2 | This is the second line.
3 | This is the third line.
```
*(Output will vary based on the actual file content)*

### Reading Specific Lines

To read only a specific range of lines (e.g., 46-68):

**Input:**
```xml
<read_file>
<path>src/app.js</path>
<start_line>46</start_line>
<end_line>68</end_line>
</read_file>
```

**Simulated Output (for lines 2-3 of `example_five_lines.txt`):**
```
2 | Content of line two.
3 | Content of line three.
```
*(Output shows only the requested lines with their original line numbers)*

### Reading a Large Text File (Automatic Truncation)

When reading a large text file without specifying a line range, the tool automatically truncates the content if it exceeds the internal line limit (e.g., 500 lines).

**Input:**
```xml
<read_file>
<path>logs/large_app.log</path>
</read_file>
```

**Simulated Output (for a 1500-line log file with a 500-line limit):**
```
1 | Log entry 1...
2 | Log entry 2...
...
500 | Log entry 500...

[Showing only 500 of 1500 total lines. Use start_line and end_line to read specific ranges.]
// Optional: Source code definitions summary might appear here for code files
```
*(Output shows the beginning lines up to the `maxReadFileLine` limit, plus a truncation notice. Use line ranges for full access.)*

### Reading Definitions Only

When `maxReadFileLine` is set to `0` in user settings, the tool returns only source code definitions without file content:

**Input:**
```xml
<!-- Assuming maxReadFileLine is set to 0 in user settings -->
<read_file>
<path>src/services/auth.service.ts</path>
</read_file>
```

**Simulated Output:**
```xml
<file>
  <path>src/services/auth.service.ts</path>
  <list_code_definition_names>
    class AuthService
      method validateUser
      method generateToken
  </list_code_definition_names>
  <notice>Showing only 0 of 150 total lines. Use line_range if you need to read more lines</notice>
</file>
```
*(This mode provides a quick overview of file structure without reading content.)*

### Attempting to Read a Non-Existent File

If the specified file does not exist:

**Input:**
```xml
<read_file>
<path>non_existent_file.txt</path>
</read_file>
```

**Simulated Output (Error):**
```
Error: File not found at path 'non_existent_file.txt'.
```

### Attempting to Read a Blocked File

If the file is excluded by rules in a `.rooignore` file:

**Input:**
```xml
<read_file>
<path>.env</path>
</read_file>
```

**Simulated Output (Error):**
```xml
<file>
  <path>.env</path>
  <error>Access denied by .rooignore rules</error>
</file>
```

---

## Image Reading Examples

The `read_file` tool now supports reading and displaying images directly in the chat interface. This enables powerful visual analysis workflows.

### Reading a Single Image

**Input:**
```xml
<read_file>
<path>assets/logo.png</path>
</read_file>
```

**Output:**
```xml
<image_content>
<path>assets/logo.png</path>
<mime_type>image/png</mime_type>
<dimensions>width: 512, height: 512</dimensions>
<data_url>data:image/png;base64,iVBORw0KGgoAAAANS...</data_url>
</image_content>
```

The image will be displayed directly in the chat interface, allowing Roo to analyze visual content.

### OCR Workflow Example

Reading multiple images from a folder for text extraction:

**Input:**
```xml
<read_file>
<args>
  <file>
    <path>screenshots/page1.png</path>
  </file>
  <file>
    <path>screenshots/page2.png</path>
  </file>
  <file>
    <path>screenshots/page3.png</path>
  </file>
</args>
</read_file>
```

**Usage:**
```
Please extract all text from these screenshot images and compile them into a single markdown document.
```

### Design Review Workflow

Analyzing multiple design mockups:

**Input:**
```xml
<read_file>
<args>
  <file>
    <path>designs/homepage-v1.jpg</path>
  </file>
  <file>
    <path>designs/homepage-v2.jpg</path>
  </file>
  <file>
    <path>designs/mobile-view.png</path>
  </file>
</args>
</read_file>
```

**Usage:**
```
Compare these design mockups and provide feedback on:
1. Visual consistency
2. Mobile responsiveness
3. Accessibility concerns
4. UI/UX improvements
```

### Supported Image Formats

The tool supports the following image formats:
- **PNG** - With dimension extraction
- **JPG/JPEG** - Standard and progressive
- **GIF** - Static and animated
- **WebP** - Modern web format
- **SVG** - Scalable vector graphics
- **BMP** - Bitmap images
- **ICO** - Icon files
- **TIFF** - Tagged image format

### Image Analysis Use Cases

1. **Documentation Screenshots**: Extract text and create documentation from UI screenshots
2. **Error Debugging**: Analyze error screenshots to understand issues
3. **Design Reviews**: Compare mockups and provide visual feedback
4. **Diagram Analysis**: Understand architecture diagrams and flowcharts
5. **Code Screenshots**: Extract code from images when text isn't available
6. **UI Testing**: Verify visual elements and layouts

---

## Multi-File Examples

When `maxConcurrentFileReads` is set to a value greater than 1, you can read multiple files simultaneously using the enhanced XML format.

### Reading Multiple Complete Files

To read several complete files at once:

**Input:**
```xml
<read_file>
<args>
  <file>
    <path>src/app.ts</path>
  </file>
  <file>
    <path>src/utils.ts</path>
  </file>
  <file>
    <path>src/config.json</path>
  </file>
</args>
</read_file>
```

**Simulated Output:**
```xml
<files>
  <file>
    <path>src/app.ts</path>
    <content>
      1 | import React from 'react'
      2 | import { Utils } from './utils'
      3 | // ... rest of file content
    </content>
  </file>
  <file>
    <path>src/utils.ts</path>
    <content>
      1 | export class Utils {
      2 |   static formatDate(date: Date): string {
      3 |     // ... utility functions
    </content>
  </file>
  <file>
    <path>src/config.json</path>
    <content>
      1 | {
      2 |   "apiUrl": "https://api.example.com",
      3 |   "timeout": 5000
      4 | }
    </content>
  </file>
</files>
```

### Reading Specific Line Ranges from Multiple Files

To read specific sections from multiple files:

**Input:**
```xml
<read_file>
<args>
  <file>
    <path>src/app.ts</path>
    <line_range>1-20</line_range>
    <line_range>45-60</line_range>
  </file>
  <file>
    <path>src/utils.ts</path>
    <line_range>10-25</line_range>
  </file>
</args>
</read_file>
```

**Simulated Output:**
```xml
<files>
  <file>
    <path>src/app.ts</path>
    <content>
      1 | import React from 'react'
      2 | import { Utils } from './utils'
      ...
      20 | const App = () => {
      
      45 |   const handleSubmit = () => {
      46 |     // Handle form submission
      ...
      60 |   }
    </content>
  </file>
  <file>
    <path>src/utils.ts</path>
    <content>
      10 |   static formatDate(date: Date): string {
      11 |     return date.toISOString().split('T')[0]
      ...
      25 |   }
    </content>
  </file>
</files>
```

### Handling Mixed Results (Some Files Denied/Blocked)

When some files are approved and others are denied or blocked:

**Input:**
```xml
<read_file>
<args>
  <file>
    <path>src/app.ts</path>
  </file>
  <file>
    <path>.env</path>
  </file>
  <file>
    <path>src/secret-config.ts</path>
  </file>
</args>
</read_file>
```

**Simulated Output:**
```xml
<files>
  <file>
    <path>src/app.ts</path>
    <content>
      1 | import React from 'react'
      2 | // ... file content successfully read
    </content>
  </file>
  <file>
    <path>.env</path>
    <error>Access denied by .rooignore rules</error>
  </file>
  <file>
    <path>src/secret-config.ts</path>
    <error>User denied access to file</error>
  </file>
</files>
```

### Batch Approval Interface

When requesting multiple files, you'll see a batch approval interface that allows you to:

- **Approve All**: Grant access to all requested files
- **Deny All**: Deny access to all requested files
- **Individual Control**: Override decisions for specific files
- **File Preview**: Click file headers to open them in your editor

The interface displays each file path clearly, making it easy to understand what Roo wants to access before granting permission.

### Mixed Content Types

You can read different types of files in a single request:

**Input:**
```xml
<read_file>
<args>
  <file>
    <path>README.md</path>
  </file>
  <file>
    <path>architecture-diagram.png</path>
  </file>
  <file>
    <path>config.json</path>
  </file>
  <file>
    <path>requirements.pdf</path>
  </file>
</args>
</read_file>
```

This allows Roo to analyze documentation, visual diagrams, configuration, and specifications all in one context.

## **search_and_replace**

---
description: Master the search_and_replace tool for finding and replacing text or regex patterns in files, with line range support and interactive diff approval.
keywords:
  - search_and_replace
  - Roo Code tools
  - find replace
  - regex patterns
  - text replacement
  - code refactoring
  - diff view
  - line ranges
  - VS Code AI
image: /img/social-share.jpg
---

# search_and_replace

The `search_and_replace` tool finds and replaces text within a file, supporting both literal strings and regular expression patterns. It allows for targeted replacements across multiple locations, optionally within specific line ranges.

---

## Parameters

### Required Parameters

- `path`: The relative path (from the workspace root) of the file to modify.
- `search`: The text string or regex pattern to find.
- `replace`: The text to replace matches with.

### Optional Parameters

- `start_line`: The 1-based line number where the search scope begins.
- `end_line`: The 1-based line number where the search scope ends (inclusive).
- `use_regex`: Set to `"true"` to treat the `search` parameter as a regular expression pattern (default is `false`).
- `ignore_case`: Set to `"true"` to perform a case-insensitive search (default is `false`).

---

## What It Does

This tool reads the specified file and performs a search-and-replace operation based on the provided parameters. It can operate on the entire file or be restricted to a specific range of lines. Changes are presented in a diff view for user review and approval before being saved.

---

## When is it used?

- When renaming variables, functions, or classes across a file.
- When updating specific text strings or values consistently.
- When applying patterned changes using regular expressions.
- When refactoring code requires replacing specific patterns.
- When making targeted changes within a defined section of a file.

---

## Key Features

- **Flexible Searching**: Supports both literal text and regular expression patterns.
- **Case Sensitivity Control**: Option to ignore case during search.
- **Scoped Replacements**: Can limit replacements to a specific range of lines (`start_line`, `end_line`).
- **Global Replacement**: Performs replacements across the entire file (or specified range) by default.
- **Interactive Approval**: Shows proposed changes in a diff view for user review and approval.
- **User Edit Support**: Allows editing the proposed content directly within the diff view.
- **Context Tracking**: Records the file edit operation.
- **Error Handling**: Checks for missing parameters, file access issues, and invalid line numbers.

---

## Limitations

- **Single File Operation**: Operates on only one file at a time. Use `search_files` to find patterns across multiple files first.
- **Review Overhead**: The mandatory diff view approval adds an interactive step.
- **Regex Complexity**: Complex regex patterns might require careful construction and testing.

---

## How It Works

When the `search_and_replace` tool is invoked, it follows this process:

1.  **Parameter Validation**: Checks for required `path`, `search`, `replace`, and validates optional parameters like line numbers and boolean flags.
2.  **File Reading**: Reads the content of the target file specified by `path`.
3.  **Regex Construction**:
    *   If `use_regex` is `false`, the `search` string is escaped to treat it as literal text.
    *   A `RegExp` object is created with the `g` (global) flag and optionally the `i` (ignore case) flag.
4.  **Replacement Execution**:
    *   If `start_line` or `end_line` are provided, the file content is split into lines, the relevant section is isolated, the replacement is performed on that section, and the file content is reconstructed.
    *   If no line range is specified, the replacement is performed on the entire file content string.
5.  **Diff View Interaction**:
    *   Opens the file in the diff view showing original vs. proposed content.
    *   Updates the diff view with the result of the replacement.
6.  **User Approval**: Presents the change via `askApproval`. Reverts if rejected.
7.  **Saving Changes**: If approved, saves the changes (including any user edits made in the diff view).
8.  **File Context Tracking**: Tracks the edit using `cline.getFileContextTracker().trackFileContext`.
9.  **Result Reporting**: Reports success (including user edits) or failure back to the AI model.

---

## Usage Examples

Simple text replacement throughout a file:

```xml
<search_and_replace>
<path>src/config.js</path>
<search>API_KEY_OLD</search>
<replace>API_KEY_NEW</replace>
</search_and_replace>
```

Case-insensitive regex replacement to update function calls:

```xml
<search_and_replace>
<path>src/app.ts</path>
<search>processData\((.*?)\)</search>
<replace>handleData($1)</replace>
<use_regex>true</use_regex>
<ignore_case>true</ignore_case>
</search_and_replace>
```

Replacing text only within lines 10 to 20:

```xml
<search_and_replace>
<path>README.md</path>
<search>Draft</search>
<replace>Final</replace>
<start_line>10</start_line>
<end_line>20</end_line>
</search_and_replace>
```

## **search_files**

---
description: Learn how search_files performs powerful regex searches across your codebase, finding patterns with context using Ripgrep for high-performance results.
keywords:
  - search_files
  - Roo Code tools
  - regex search
  - code patterns
  - Ripgrep
  - multi-file search
  - codebase search
  - pattern matching
  - VS Code AI
image: /img/social-share.jpg
---

# search_files

The `search_files` tool performs regex searches across multiple files within your project's workspace. For security, it cannot search outside the current workspace directory. It helps Roo locate specific code patterns, text, or other content throughout your codebase with contextual results.

---

## Parameters

The tool accepts these parameters:

- `path` (required): The path of the directory to search in, relative to the current workspace directory. The search is confined to the workspace.
- `regex` (required): The regular expression pattern to search for (uses Rust regex syntax)
- `file_pattern` (optional): Glob pattern to filter files (e.g., '*.ts' for TypeScript files)

---

## What It Does

This tool searches across files in a specified directory using regular expressions, showing each match with surrounding context. It's like having a powerful "Find in Files" feature that works across the entire project structure.

---

## When is it used?

- When Roo needs to find where specific functions or variables are used
- When Roo helps with refactoring and needs to understand usage patterns
- When Roo needs to locate all instances of a particular code pattern
- When Roo searches for text across multiple files with filtering capabilities

---

## Key Features

- Searches across multiple files in a single operation using high-performance Ripgrep
- Shows context around each match (1 line before and after)
- Filters files by type using glob patterns (e.g., only TypeScript files)
- Provides line numbers for easy reference
- Uses powerful regex patterns for precise searches
- Automatically limits output to 300 results with notification
- Truncates lines longer than 500 characters with "[truncated...]" marker
- Intelligently combines nearby matches into single blocks for readability

---

## Limitations

- Works best with text-based files (not effective for binary files like images)
- Performance may slow with extremely large codebases
- Uses Rust regex syntax, which may differ slightly from other regex implementations
- Cannot search within compressed files or archives
- Default context size is fixed (1 line before and after)
- May display varying context sizes when matches are close together due to result grouping
- For security, searches are strictly limited to the current workspace and cannot access parent directories or other locations on the file system.

---

## How It Works

When the `search_files` tool is invoked, it follows this process:

1. **Parameter Validation**: Validates the required `path` and `regex` parameters
2. **Path Resolution**: Resolves the relative path to an absolute path
3. **Search Execution**:
   - Uses Ripgrep (rg) for high-performance text searching
   - Applies file pattern filtering if specified
   - Collects matches with surrounding context
4. **Result Formatting**:
   - Formats results with file paths, line numbers, and context
   - Displays 1 line of context before and after each match
   - Structures output for easy readability
   - Limits results to a maximum of 300 matches with notification
   - Truncates lines longer than 500 characters
   - Merges nearby matches into contiguous blocks

---

## Search Results Format

The search results include:

- Relative file paths for each matching file (prefixed with #)
- Context lines before and after each match (1 line by default)
- Line numbers padded to 3 spaces followed by ` | ` and the line content
- A separator line (----) after each match group

Example output format:
```
# rel/path/to/app.ts
 11 |   // Some processing logic here
 12 |   // TODO: Implement error handling
 13 |   return processedData;
----

# Showing first 300 of 300+ results. Use a more specific search if necessary.
```

When matches occur close to each other, they're merged into a single block rather than shown as separate results:

```
# rel/path/to/auth.ts
 13 | // Some code here
 14 | // TODO: Add proper validation
 15 | function validateUser(credentials) {
 16 |   // TODO: Implement rate limiting
 17 |   return checkDatabase(credentials);
----
```

---

## Examples When Used

- When asked to refactor a function, Roo first searches for all places the function is used to ensure comprehensive changes.
- When investigating bugs, Roo searches for similar patterns to identify related issues across the codebase.
- When addressing technical debt, Roo locates all TODO comments across the project.
- When analyzing dependencies, Roo finds all imports of a particular module.

---

## Usage Examples

Searching for TODO comments in all JavaScript files:
```
<search_files>
<path>src</path>
<regex>TODO|FIXME</regex>
<file_pattern>*.js</file_pattern>
</search_files>
```

Finding all usages of a specific function:
```
<search_files>
<path>.</path>
<regex>function\s+calculateTotal</regex>
<file_pattern>*.{js,ts}</file_pattern>
</search_files>
```

Searching for a specific import pattern across the entire project:
```
<search_files>
<path>.</path>
<regex>import\s+.*\s+from\s+['"]@components/</regex>
</search_files>
```

## **switch_mode**

---
description: Understand how switch_mode enables seamless transitions between Code, Architect, Ask, and Debug modes for specialized task handling in Roo Code.
keywords:
  - switch_mode
  - Roo Code tools
  - mode switching
  - operational modes
  - Code mode
  - Architect mode
  - Debug mode
  - Ask mode
  - task transitions
  - VS Code AI
image: /img/social-share.jpg
---

# switch_mode

The `switch_mode` tool enables Roo to change between different operational modes, each with specialized capabilities for specific types of tasks. This allows seamless transitions between modes like Code, Architect, Ask, or Debug when the current task requires different expertise.

---

## Parameters

The tool accepts these parameters:

- `mode_slug` (required): The slug of the mode to switch to (e.g., "code", "ask", "architect")
- `reason` (optional): The reason for switching modes, providing context for the user

---

## What It Does

This tool requests a mode change when the current task would be better handled by another mode's capabilities. It maintains context while shifting Roo's focus and available toolsets to match the requirements of the new task phase.

---

## When is it used?

- When transitioning from information gathering to code implementation
- When shifting from coding to architecture or design
- When the current task requires capabilities only available in a different mode
- When specialized expertise is needed for a particular phase of a complex project

---

## Key Features

- Maintains context continuity across mode transitions
- Provides clear reasoning for mode switch recommendations
- Requires user approval for all mode changes
- Enforces tool group restrictions specific to each mode
- Seamlessly adapts tool availability based on the selected mode
- Works with both standard and custom modes
- Displays the mode switch and reasoning in the UI
- Uses XML-style formatting for parameter specification
- Handles file type restrictions specific to certain modes

---

## Limitations

- Cannot switch to modes that don't exist in the system
- Requires explicit user approval for each mode transition
- Cannot use tools specific to a mode until the switch is complete
- Applies a 500ms delay after mode switching to allow the change to take effect
- Some modes have file type restrictions (e.g., Architect mode can only edit markdown files)
- Mode preservation for resumption applies only to the `new_task` functionality, not general mode switching

---

## How It Works

When the `switch_mode` tool is invoked, it follows this process:

1. **Request Validation**:
   - Validates that the requested mode exists in the system
   - Checks that the `mode_slug` parameter is provided and valid
   - Verifies the user isn't already in the requested mode
   - Ensures the `reason` parameter (if provided) is properly formatted

2. **Mode Transition Preparation**:
   - Packages the mode change request with the provided reason
   - Presents the change request to the user for approval

3. **Mode Activation (Upon User Approval)**:
   - Updates the UI to reflect the new mode
   - Adjusts available tools based on the mode's tool group configuration
   - Applies the mode-specific prompt and behavior
   - Applies a 500ms delay to allow the change to take effect before executing next tool
   - Enforces any file restrictions specific to the mode

4. **Continuation**:
   - Proceeds with the task using the capabilities of the new mode
   - Retains relevant context from the previous interaction

---

## Tool Group Association

The `switch_mode` tool belongs to the "modes" tool group but is also included in the "always available" tools list. This means:

- It can be used in any mode regardless of the mode's configured tool groups
- It's available alongside other core tools like `ask_followup_question` and `attempt_completion`
- It allows mode transitions at any point in a workflow when task requirements change

---

## Mode Structure

Each mode in the system has a specific structure:

- `slug`: Unique identifier for the mode (e.g., "code", "ask")
- `name`: Display name for the mode (e.g., "Code", "Ask")
- `roleDefinition`: The specialized role and capabilities of the mode
- `customInstructions`: Optional mode-specific instructions that guide behavior
- `groups`: Tool groups available to the mode with optional restrictions

---

## Mode Capabilities

The core modes provide these specialized capabilities:

- **Code Mode**: Focused on coding tasks with full access to code editing tools
- **Architect Mode**: Specialized for system design and architecture planning, limited to editing markdown files only
- **Ask Mode**: Optimized for answering questions and providing information
- **Debug Mode**: Equipped for systematic problem diagnosis and resolution

---

## Custom Modes

Beyond the core modes, the system supports custom project-specific modes:

- Custom modes can be defined with specific tool groups enabled
- They can specify custom role definitions and instructions
- The system checks custom modes first before falling back to core modes
- Custom mode definitions take precedence over core modes with the same slug

---

## File Restrictions

Different modes may have specific file type restrictions:

- **Architect Mode**: Can only edit files matching the `.md` extension
- Attempting to edit restricted file types results in a `FileRestrictionError`
- These restrictions help enforce proper separation of concerns between modes

---

## Examples When Used

- When discussing a new feature, Roo switches from Ask mode to Architect mode to help design the system structure.
- After completing architecture planning in Architect mode, Roo switches to Code mode to implement the designed features.
- When encountering bugs during development, Roo switches from Code mode to Debug mode for systematic troubleshooting.

---

## Usage Examples

Switching to Code mode for implementation:
```
<switch_mode>
<mode_slug>code</mode_slug>
<reason>Need to implement the login functionality based on the architecture we've discussed</reason>
</switch_mode>
```

Switching to Architect mode for design:
```
<switch_mode>
<mode_slug>architect</mode_slug>
<reason>Need to design the system architecture before implementation</reason>
</switch_mode>
```

Switching to Debug mode for troubleshooting:
```
<switch_mode>
<mode_slug>debug</mode_slug>
<reason>Need to systematically diagnose the authentication error</reason>
</switch_mode>
```

Switching to Ask mode for information:
```
<switch_mode>
<mode_slug>ask</mode_slug>
<reason>Need to answer questions about the implemented feature</reason>
</switch_mode>
```

## **Tool Use Overview**

---
description: Comprehensive guide to Roo Code's tool system, including tool groups, calling mechanisms, mode integration, and best practices for AI-powered development.
keywords:
  - Roo Code tools
  - tool system
  - tool groups
  - AI development
  - tool architecture
  - mode integration
  - tool security
  - workflow tools
  - VS Code AI
image: /img/social-share.jpg
---

# Tool Use Overview

Roo Code implements a sophisticated tool system that allows AI models to interact with your development environment in a controlled and secure manner. This document explains how tools work, when they're called, and how they're managed.

---

## Core Concepts

### Tool Groups

Tools are organized into logical groups based on their functionality:

| Category | Purpose | Tools | Common Use |
|----------|---------|-------|------------|
| **Read Group** | File system reading and exploration | [read_file](/advanced-usage/available-tools/read-file), [list_files](/advanced-usage/available-tools/list-files), [list_code_definition_names](/advanced-usage/available-tools/list-code-definition-names) | Code exploration and analysis |
| **Search Group** | Pattern and semantic searching | [search_files](/advanced-usage/available-tools/search-files), [codebase_search](/advanced-usage/available-tools/codebase-search) | Finding code patterns and functionality |
| **Edit Group** | File system modifications | [apply_diff](/advanced-usage/available-tools/apply-diff), [insert_content](/advanced-usage/available-tools/insert-content), [search_and_replace](/advanced-usage/available-tools/search-and-replace), [write_to_file](/advanced-usage/available-tools/write-to-file) | Code changes and file manipulation |
| **Browser Group** | Web automation | [browser_action](/advanced-usage/available-tools/browser-action) | Web testing and interaction |
| **Command Group** | System command execution | [execute_command](/advanced-usage/available-tools/execute-command) | Running scripts, building projects |
| **MCP Group** | External tool integration | [use_mcp_tool](/advanced-usage/available-tools/use-mcp-tool), [access_mcp_resource](/advanced-usage/available-tools/access-mcp-resource) | Specialized functionality through external servers |
| **Workflow Group** | Mode and task management | [switch_mode](/advanced-usage/available-tools/switch-mode), [new_task](/advanced-usage/available-tools/new-task), [ask_followup_question](/advanced-usage/available-tools/ask-followup-question), [attempt_completion](/advanced-usage/available-tools/attempt-completion) | Context switching and task organization |

### Always Available Tools

Certain tools are accessible regardless of the current mode:

- [ask_followup_question](/advanced-usage/available-tools/ask-followup-question): Gather additional information from users
- [attempt_completion](/advanced-usage/available-tools/attempt-completion): Signal task completion
- [switch_mode](/advanced-usage/available-tools/switch-mode): Change operational modes
- [new_task](/advanced-usage/available-tools/new-task): Create subtasks

---

## Available Tools

### Read Tools
These tools help Roo understand your code and project:

- [read_file](/advanced-usage/available-tools/read-file) - Examines the contents of files
- [list_files](/advanced-usage/available-tools/list-files) - Maps your project's file structure
- [list_code_definition_names](/advanced-usage/available-tools/list-code-definition-names) - Creates a structural map of your code

### Search Tools
These tools help Roo find patterns and functionality across your codebase:

- [search_files](/advanced-usage/available-tools/search-files) - Finds patterns across multiple files using regex
- [codebase_search](/advanced-usage/available-tools/codebase-search) - Performs semantic searches across your indexed codebase

### Edit Tools
These tools help Roo make changes to your code:

- [apply_diff](/advanced-usage/available-tools/apply-diff) - Makes precise, surgical changes to your code
- [insert_content](/advanced-usage/available-tools/insert-content) - Adds new lines of content without modifying existing lines
- [search_and_replace](/advanced-usage/available-tools/search-and-replace) - Finds and replaces text or regex patterns within a file
- [write_to_file](/advanced-usage/available-tools/write-to-file) - Creates new files or completely rewrites existing ones

### Browser Tools
These tools help Roo interact with web applications:

- [browser_action](/advanced-usage/available-tools/browser-action) - Automates browser interactions

### Command Tools
These tools help Roo execute commands:

- [execute_command](/advanced-usage/available-tools/execute-command) - Runs system commands and programs

### MCP Tools
These tools help Roo connect with external services:

- [use_mcp_tool](/advanced-usage/available-tools/use-mcp-tool) - Uses specialized external tools
- [access_mcp_resource](/advanced-usage/available-tools/access-mcp-resource) - Accesses external data sources

### Workflow Tools
These tools help manage the conversation and task flow:

- [ask_followup_question](/advanced-usage/available-tools/ask-followup-question) - Gets additional information from you
- [attempt_completion](/advanced-usage/available-tools/attempt-completion) - Presents final results
- [switch_mode](/advanced-usage/available-tools/switch-mode) - Changes to a different mode for specialized tasks
- [new_task](/advanced-usage/available-tools/new-task) - Creates a new subtask

---

## Tool Calling Mechanism

### Handling Complex Tasks

For certain complex operations that require multiple steps, Roo doesn't just figure them out on the fly. Instead, it follows predefined, internal plans to ensure consistency and accuracy.

A prime example is creating a new MCP server, identified internally by `create_mcp_server`. **This identifier does not represent a tool you will see being called.** Rather, when you ask Roo to create a server, it triggers this known, multi-step workflow.

This specific workflow is initiated by Roo using its internal `fetch_instructions` tool (with the task `create_mcp_server`) to retrieve a detailed plan. This plan then guides Roo to make calls to several standard, documented tools in sequence, such as:

*   [`execute_command`](/advanced-usage/available-tools/execute-command) for running setup scripts (e.g., `npx @modelcontextprotocol/create-server`).
*   [`write_to_file`](/advanced-usage/available-tools/write-to-file) or [`apply_diff`](/advanced-usage/available-tools/apply-diff) for creating or modifying server code and configuration files.
*   [`ask_followup_question`](/advanced-usage/available-tools/ask-followup-question) to gather necessary information like API keys from you.
*   Other standard tools as needed for steps like determining file locations or updating configuration entries.

So, while the overall task (like `create_mcp_server`) is complex, it's ultimately accomplished by intelligently orchestrating the standard tools available in your environment. This approach allows Roo to reliably perform complex operations by leveraging the tools documented here.

### When Tools Are Called

Tools are invoked under specific conditions:

1. **Direct Task Requirements**
   - When specific actions are needed to complete a task as decided by the LLM
   - In response to user requests
   - During automated workflows

2. **Mode-Based Availability**
   - Different modes enable different tool sets
   - Mode switches can trigger tool availability changes
   - Some tools are restricted to specific modes

3. **Context-Dependent Calls**
   - Based on the current state of the workspace
   - In response to system events
   - During error handling and recovery

### Decision Process

The system uses a multi-step process to determine tool availability:

1. **Mode Validation**
   ```typescript
   isToolAllowedForMode(
       tool: string,
       modeSlug: string,
       customModes: ModeConfig[],
       toolRequirements?: Record<string, boolean>,
       toolParams?: Record<string, any>
   )
   ```

2. **Requirement Checking**
   - System capability verification
   - Resource availability
   - Permission validation

3. **Parameter Validation**
   - Required parameter presence
   - Parameter type checking
   - Value validation

---

## Technical Implementation

### Tool Call Processing

1. **Initialization**
   - Tool name and parameters are validated
   - Mode compatibility is checked
   - Requirements are verified

2. **Execution**
   ```typescript
   const toolCall = {
       type: "tool_call",
       name: chunk.name,
       arguments: chunk.input,
       callId: chunk.callId
   }
   ```

3. **Result Handling**
   - Success/failure determination
   - Result formatting
   - Error handling

### Security and Permissions

1. **Access Control**
   - File system restrictions
   - Command execution limitations
   - Network access controls

2. **Validation Layers**
   - Tool-specific validation
   - Mode-based restrictions
   - System-level checks

---

## Mode Integration

### Mode-Based Tool Access

Tools are made available based on the current mode:

- **Code Mode**: Full access to file system tools, code editing capabilities, command execution
- **Ask Mode**: Limited to reading tools, information gathering capabilities, no file system modifications
- **Architect Mode**: Design-focused tools, documentation capabilities, limited execution rights
- **Custom Modes**: Can be configured with specific tool access for specialized workflows

### Mode Switching

1. **Process**
   - Current mode state preservation
   - Tool availability updates
   - Context switching

2. **Impact on Tools**
   - Tool set changes
   - Permission adjustments
   - Context preservation

---

## Best Practices

### Tool Usage Guidelines

1. **Efficiency**
   - Use the most specific tool for the task
   - Avoid redundant tool calls
   - Batch operations when possible

2. **Security**
   - Validate inputs before tool calls
   - Use minimum required permissions
   - Follow security best practices

3. **Error Handling**
   - Implement proper error checking
   - Provide meaningful error messages
   - Handle failures gracefully

### Common Patterns

1. **Information Gathering**
   ```
   [ask_followup_question](/advanced-usage/available-tools/ask-followup-question) → [read_file](/advanced-usage/available-tools/read-file) → [codebase_search](/advanced-usage/available-tools/codebase-search)
   ```

2. **Code Modification**
   ```
   [read_file](/advanced-usage/available-tools/read-file) → [apply_diff](/advanced-usage/available-tools/apply-diff) → [attempt_completion](/advanced-usage/available-tools/attempt-completion)
   ```

3. **Task Management**
   ```
   [new_task](/advanced-usage/available-tools/new-task) → [switch_mode](/advanced-usage/available-tools/switch-mode) → [execute_command](/advanced-usage/available-tools/execute-command)
   ```

---

## Error Handling and Recovery

### Error Types

1. **Tool-Specific Errors**
   - Parameter validation failures
   - Execution errors
   - Resource access issues

2. **System Errors**
   - Permission denied
   - Resource unavailable
   - Network failures

3. **Context Errors**
   - Invalid mode for tool
   - Missing requirements
   - State inconsistencies

### Recovery Strategies

1. **Automatic Recovery**
   - Retry mechanisms
   - Fallback options
   - State restoration

2. **User Intervention**
   - Error notifications
   - Recovery suggestions
   - Manual intervention options

## **update_todo_list**

---
description: Learn how update_todo_list creates dynamic TODO lists with status tracking, enabling step-by-step task management for complex workflows in Roo Code.
keywords:
  - update_todo_list
  - Roo Code tools
  - task management
  - TODO lists
  - workflow tracking
  - checklist management
  - task status
  - interactive UI
  - VS Code AI
image: /img/social-share.jpg
---

# update_todo_list

The `update_todo_list` tool enables dynamic, interactive task management within the chat interface. It replaces the entire TODO list with an updated checklist, ensuring that task status is always current and providing step-by-step tracking for complex, multi-step workflows.

---

## Parameters

The tool accepts these parameters:

- `todos` (required): A markdown-formatted string representing the complete checklist with status indicators

---

## What It Does

This tool creates and manages an interactive todo list that appears as a UI component in the chat interface. It allows for real-time task tracking, status updates, and dynamic addition of new items as they are discovered during complex workflows. The list provides a structured way to manage multi-step tasks with clear visual progress indicators.

---

## When is it used?

- When managing complex, multi-step tasks that benefit from structured tracking
- When Roo needs to show progress through a series of related activities
- When tasks require step-by-step completion verification before proceeding
- When new actionable items are discovered during long or complex workflows
- When providing clear checkpoints and progress visibility to users

---

## Key Features

- **Full Checklist Replacement**: Overwrites the existing todo list with the updated version provided
- **Interactive UI Component**: Displays as an editable interface element in the chat
- **Multiple Status Types**: Supports pending, in-progress, and completed task states
- **Dynamic Task Management**: Add new tasks as they arise during workflow execution
- **User-Friendly Editing**: Provides direct editing capabilities within the chat interface
- **Step-by-Step Tracking**: Enables confirmation of each step before updating and proceeding
- **Progress Visualization**: Clear visual indicators for task completion status
- **Workflow Integration**: Seamlessly integrates with task execution and completion flows

---

## Limitations

- **Complete Replacement**: Replaces the entire list rather than making incremental updates
- **Single-Level Structure**: Uses single-level markdown checklists without nesting support
- **Format Requirements**: Requires specific markdown checkbox syntax for proper parsing
- **Manual Updates**: Requires explicit tool calls to update the list status
- **State Management**: Todo list state is tied to the current task and conversation context

---

## How It Works

When the `update_todo_list` tool is invoked, it follows this process:

1. **Input Validation**: 
   - Validates the required `todos` parameter is provided
   - Parses the markdown checklist format for syntax correctness
   - Checks for valid status indicators: `[ ]`, `[-]`, and `[x]`

2. **List Processing**:
   - Processes the markdown-formatted checklist
   - Extracts individual todo items with their status indicators
   - Validates the structure and format of each item

3. **UI Integration**:
   - Presents the updated todo list to the user for approval
   - Replaces any existing todo list with the new version
   - Renders the list as an interactive component in the chat interface

4. **User Interaction**:
   - Allows users to edit todos directly in the UI when in editing mode
   - Provides "Add Todo" functionality for real-time list expansion
   - Synchronizes changes back to the extension to maintain state consistency

5. **State Management**:
   - Updates the task's internal todo list representation
   - Maintains synchronization between UI state and backend data
   - Preserves todo list state across conversation interactions

---

## Checklist Format Requirements

The tool uses a specific markdown format for todo items:

### Status Options
- `[ ]` - Pending task (not started)
- `[-]` - In progress task (currently being worked on)  
- `[x]` - Completed task (fully finished)

### Format Rules
- Use single-level markdown checklist (no nesting or subtasks)
- List todos in intended execution order
- Each todo item should be clear and actionable
- Status should accurately reflect current task state

---

## Task Management Guidelines

### Status Updates
- Mark tasks as completed immediately after all work is finished
- Start the next task by marking it as in progress
- Use pending status for tasks not yet started
- Only mark tasks as completed when fully accomplished with no unresolved dependencies

### Dynamic List Management
- Add new todos as soon as they are identified during task execution
- Remove tasks only if they are no longer relevant or explicitly requested
- Retain all unfinished tasks and update their status as needed
- If a task is blocked, keep it as in progress and add new todos for resolution steps

---

## Examples When Used

- When developing a web application, Roo creates a todo list tracking design, implementation, testing, and deployment phases.
- When setting up a development environment, Roo tracks installation of dependencies, configuration steps, and verification tasks.
- When debugging complex issues, Roo maintains a list of investigation steps, potential causes, and testing procedures.
- When refactoring code, Roo tracks which files need updates, what tests need modification, and documentation changes required.
- When implementing new features, Roo manages tasks for planning, coding, testing, and integration steps.

---

## Usage Examples

Creating an initial todo list for a development task:
```xml
<update_todo_list>
<todos>
[ ] Analyze requirements
[ ] Design architecture  
[ ] Implement core logic
[ ] Write tests
[ ] Update documentation
</todos>
</update_todo_list>
```

Updating progress after completing the first task:
```xml
<update_todo_list>
<todos>
[x] Analyze requirements
[-] Design architecture
[ ] Implement core logic
[ ] Write tests
[ ] Update documentation
</todos>
</update_todo_list>
```

Adding new tasks discovered during implementation:
```xml
<update_todo_list>
<todos>
[x] Analyze requirements
[x] Design architecture
[x] Implement core logic
[-] Write tests
[ ] Update documentation
[ ] Add performance benchmarks
[ ] Create deployment script
</todos>
</update_todo_list>
```

Managing a complex debugging workflow:
```xml
<update_todo_list>
<todos>
[x] Reproduce the issue
[x] Check recent code changes
[-] Analyze error logs
[ ] Test with different configurations
[ ] Check database queries
[ ] Verify network connectivity
[ ] Create fix and test
</todos>
</update_todo_list>
```

## **use_mcp_tool**

---
description: Discover how use_mcp_tool integrates external MCP servers to extend Roo Code with specialized tools, domain-specific functionality, and external services.
keywords:
  - use_mcp_tool
  - MCP tools
  - Model Context Protocol
  - external tools
  - Roo Code integration
  - MCP servers
  - domain tools
  - tool extension
  - VS Code AI
image: /img/social-share.jpg
---

# use_mcp_tool

The `use_mcp_tool` tool enables interaction with external tools provided by connected Model Context Protocol (MCP) servers. It extends Roo's capabilities with domain-specific functionality through a standardized protocol.

---

## Parameters

The tool accepts these parameters:

- `server_name` (required): The name of the MCP server providing the tool
- `tool_name` (required): The name of the tool to execute
- `arguments` (required/optional): A JSON object containing the tool's input parameters, following the tool's input schema. May be optional for tools that require no input.

---

## What It Does

This tool allows Roo to access specialized functionality provided by external MCP servers. Each MCP server can offer multiple tools with unique capabilities, extending Roo beyond its built-in functionality. The system validates arguments against schemas, manages server connections, and processes responses of various content types (text, image, resource).

---

## When is it used?

- When specialized functionality not available in core tools is needed
- When domain-specific operations are required
- When integration with external systems or services is needed
- When working with data that requires specific processing or analysis
- When accessing proprietary tools through a standardized interface

---

## Key Features

- Uses the standardized MCP protocol via the `@modelcontextprotocol/sdk` library
- Supports multiple transport mechanisms (StdioClientTransport, StreamableHTTPClientTransport and SSEClientTransport)
- Validates arguments using Zod schema validation on both client and server sides
- Processes multiple response content types: text, image, and resource references
- Manages server lifecycle with automatic restarts when server code changes
- Provides an "always allow" mechanism to bypass approval for trusted tools
- Works with the companion `access_mcp_resource` tool for resource retrieval
- Maintains proper error tracking and handling for failed operations
- Supports configurable timeouts (1-3600 seconds, default: 60 seconds)
- Allows file watchers to automatically detect and reload server changes

---

## Limitations

- Depends on external MCP servers being available and connected
- Limited to the tools provided by connected servers
- Tool capabilities vary between different MCP servers
- Network issues can affect reliability and performance
- Requires user approval before execution (unless in the "always allow" list)
- Cannot execute multiple MCP tool operations simultaneously

---

## Server Configuration

MCP servers can be configured globally or at the project level:

- **Global Configuration**: Managed through the Roo Code extension settings in VS Code. These apply across all projects unless overridden.
- **Project-level Configuration**: Defined in a `.roo/mcp.json` file within your project's root directory.
 - This allows project-specific server setups.
 - Project-level servers take precedence over global servers if they share the same name.
 - Since `.roo/mcp.json` can be committed to version control, it simplifies sharing configurations with your team.

---

## How It Works

When the `use_mcp_tool` tool is invoked, it follows this process:

1. **Initialization and Validation**:
   - The system verifies that the MCP hub is available
   - Confirms the specified server exists and is connected
   - Validates the requested tool exists on the server
   - Arguments are validated against the tool's schema definition
   - Timeout settings are extracted from server configuration (default: 60 seconds)

2. **Execution and Communication**:
   - The system selects the appropriate transport mechanism:
     - `StdioClientTransport`: For communicating with local processes via standard I/O
     - `SSEClientTransport`: For communicating with HTTP servers via Server-Sent Events
     - `StreamableHTTPClientTransport`: For communicating with HTTP servers via Streamable HTTP Events
   - A request is sent with validated server name, tool name, and arguments
   - Communication uses the `@modelcontextprotocol/sdk` library for standardized interactions
   - Request execution is tracked with timeout handling to prevent hanging operations

3. **Response Processing**:
   - Responses can include multiple content types:
     - Text content: Plain text responses
     - Image content: Binary image data with MIME type information
     - Resource references: URIs to access server resources (works with `access_mcp_resource`)
   - The system checks the `isError` flag to determine if error handling is needed
   - Results are formatted for display in the Roo interface

4. **Resource and Error Handling**:
   - The system uses WeakRef patterns to prevent memory leaks
   - A consecutive mistake counter tracks and manages errors
   - File watchers monitor for server code changes and trigger automatic restarts
   - The security model requires approval for tool execution unless in the "always allow" list

---

## Security and Permissions

The MCP architecture provides several security features:

- Users must approve tool usage before execution (by default)
- Specific tools can be marked for automatic approval in the "always allow" list
- Server configurations are validated with Zod schemas for integrity
- Configurable timeouts prevent hanging operations (1-3600 seconds)
- Server connections can be enabled or disabled through the UI

---

## Examples When Used

- Analyzing specialized data formats using server-side processing tools
- Generating images or other media through AI models hosted on external servers
- Executing complex domain-specific calculations without local implementation
- Accessing proprietary APIs or services through a controlled interface
- Retrieving data from specialized databases or data sources

---

## Usage Examples

Requesting weather forecast data with text response:
```
<use_mcp_tool>
<server_name>weather-server</server_name>
<tool_name>get_forecast</tool_name>
<arguments>
{
  "city": "San Francisco",
  "days": 5,
  "format": "text"
}
</arguments>
</use_mcp_tool>
```

Analyzing source code with a specialized tool that returns JSON:
```
<use_mcp_tool>
<server_name>code-analysis</server_name>
<tool_name>complexity_metrics</tool_name>
<arguments>
{
  "language": "typescript",
  "file_path": "src/app.ts",
  "include_functions": true,
  "metrics": ["cyclomatic", "cognitive"]
}
</arguments>
</use_mcp_tool>
```

Generating an image with specific parameters:
```
<use_mcp_tool>
<server_name>image-generation</server_name>
<tool_name>create_image</tool_name>
<arguments>
{
  "prompt": "A futuristic city with flying cars",
  "style": "photorealistic",
  "dimensions": {
    "width": 1024,
    "height": 768
  },
  "format": "webp"
}
</arguments>
</use_mcp_tool>
```

Accessing a resource through a tool that returns a resource reference:
```
<use_mcp_tool>
<server_name>database-connector</server_name>
<tool_name>query_and_store</tool_name>
<arguments>
{
  "database": "users",
  "type": "select",
  "fields": ["name", "email", "last_login"],
  "where": {
    "status": "active"
  },
  "store_as": "active_users"
}
</arguments>
</use_mcp_tool>
```

Tool with no required arguments:
```
<use_mcp_tool>
<server_name>system-monitor</server_name>
<tool_name>get_current_status</tool_name>
<arguments>
{}
</arguments>
</use_mcp_tool>
```

## **write_to_file**

---
description: Explore write_to_file for creating new files or replacing content with interactive diff view approval, ensuring safe file operations in Roo Code.
keywords:
  - write_to_file
  - Roo Code tools
  - file creation
  - file writing
  - diff view
  - content approval
  - file operations
  - interactive editing
  - VS Code AI
image: /img/social-share.jpg
---

# write_to_file

The `write_to_file` tool creates new files or completely replaces existing file content with an interactive approval process. It provides a diff view for reviewing changes before they're applied.

---

## Parameters

The tool accepts these parameters:

- `path` (required): The path of the file to write to, relative to the current working directory
- `content` (required): The complete content to write to the file
- `line_count` (required): The number of lines in the file, including empty lines

---

## What It Does

This tool writes content to a specified file, either creating a new file if it doesn't exist or completely overwriting an existing file. All changes require explicit user approval through a diff view interface, where users can review and even edit the proposed changes before they're applied.

---

## When is it used?

- When Roo needs to create a new file from scratch
- When Roo needs to completely rewrite an existing file
- When creating multiple files for a new project
- When generating configuration files, documentation, or source code
- When you need to review changes before they're applied

---

## Key Features

- Interactive Approval: Shows changes in a diff view requiring explicit approval before applying
- User Edit Support: Allows editing the proposed content before final approval
- Safety Measures: Detects code omission, validates paths, and prevents truncated content
- Editor Integration: Opens a diff view that scrolls to the first difference automatically
- Content Preprocessing: Handles artifacts from different AI models to ensure clean content
- Access Control: Validates against `.rooignore` restrictions before making changes
- Parent Directories: May handle directory creation through system dependencies
- Complete Replacement: Provides a fully transformed file in a single operation

---

## Limitations

- Not suitable for existing files: Much slower and less efficient than `apply_diff` for modifying existing files
- Performance with large files: Operation becomes significantly slower with larger files
- Complete overwrite: Replaces entire file content, cannot preserve original content
- Line count required: Needs accurate line count to detect potential content truncation
- Review overhead: The approval process adds extra steps compared to direct edits
- Interactive only: Cannot be used in automated workflows that require non-interactive execution

---

## How It Works

When the `write_to_file` tool is invoked, it follows this process:

1. **Parameter Validation**: Validates the required parameters and permissions
   - Checks that `path`, `content`, and `line_count` are provided
   - If `line_count` is missing/invalid, reverts any diff view changes and returns an error suggesting alternative tools (`apply_diff`, `insert_content`, etc.) if modifying an existing file.
   - Validates the file is allowed (not restricted by `.rooignore`)
   - Ensures the path is within the workspace boundaries
   - Tracks consecutive mistake counts for missing parameters
   - Shows specific error messages for each validation failure

2. **Content Preprocessing**:
   - Removes code block markers that might be added by AI models
   - Handles escaped HTML entities (specifically for non-Claude models)
   - Strips line numbers if accidentally included in content
   - Performs model-specific processing for different AI providers

3. **Diff View Generation**:
   - Opens a diff view in the editor showing the proposed changes
   - Adds a 300ms delay to ensure UI responsiveness
   - Scrolls automatically to the first difference
   - Highlights changes for easy review

4. **User Approval Process**:
   - Waits for explicit user approval to proceed
   - Allows users to edit the content in the diff view
   - Captures any user edits for the final content
   - Provides option to reject changes entirely
   - Detects and incorporates user modifications into the final result

5. **Safety Validation**:
   - Detects potential content truncation by comparing with provided line count
   - Shows warnings if content appears incomplete
   - Validates file path and access permissions
   - Specifically checks if files are outside the workspace with `isOutsideWorkspace` flag

6. **File Writing**:
   - Writes the approved content (with any user edits) to the file
   - Provides confirmation of successful write
   - Resets the consecutive mistakes counter on success

---

## Examples When Used

- When creating a new project, Roo generates multiple files but lets you review each before committing changes.
- When setting up configuration files, Roo shows the proposed configuration in a diff view for approval.
- When generating documentation, Roo creates markdown files but lets you make final adjustments in the diff view.
- When developing a prototype, Roo shows complete source files in a diff view where you can fine-tune before saving.

---

## Usage Examples

Creating a new JSON configuration file:
```
<write_to_file>
<path>config/settings.json</path>
<content>
{
  "apiEndpoint": "https://api.example.com",
  "theme": {
    "primaryColor": "#007bff",
    "secondaryColor": "#6c757d",
    "fontFamily": "Arial, sans-serif"
  },
  "features": {
    "darkMode": true,
    "notifications": true,
    "analytics": false
  },
  "version": "1.0.0"
}
</content>
<line_count>14</line_count>
</write_to_file>
```

Creating a simple HTML file:
```
<write_to_file>
<path>src/index.html</path>
<content>
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>My Application</title>
  <link rel="stylesheet" href="styles.css">
</head>
<body>
  <div id="app"></div>
  <script src="app.js"></script>
</body>
</html>
</content>
<line_count>13</line_count>
</write_to_file>
```

Creating a JavaScript module:
```
<write_to_file>
<path>src/utils/helpers.js</path>
<content>
/**
 * Utility functions for the application
 */

export function formatDate(date) {
  return new Date(date).toLocaleDateString();
}

export function calculateTotal(items) {
  return items.reduce((sum, item) => sum + item.price, 0);
}

export function debounce(func, delay) {
  let timeout;
  return function(...args) {
    clearTimeout(timeout);
    timeout = setTimeout(() => func.apply(this, args), delay);
  };
}
</content>
<line_count>18</line_count>
</write_to_file>
```