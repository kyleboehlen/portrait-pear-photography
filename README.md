# Portrait Pear Portfolio

This is my photography portfolio application designed to showcase my work with minimal cloud hosting costs. I chose <br>
GoLang for the API because it's straightforward, portable, includes an HTTP server in the standard library, and has <br>
fast cold starts.

The frontend uses Vue3 with JavaScript. This builds on my original Portrait Pear app, which was messy and included <br>
Capacitor for iOS support and offline capabilities. I've cleaned up the code, removed unnecessary features (including <br>
Capacitor), updated dependencies, and added an API.

While I'd make different technology choices if starting fresh (TypeScript for type safety, Vitest for testing, and <br>
avoiding DaisyUI's limitations, FSD anyone?), this is a small, single-purpose app. Rather than succumbing to the<br>
greenfield temptation, I'm evolving the existing codebase.

### New additions:
- pnpm package management
- GoLang API
- Cloudflare Images for hosting and resizing
- Cloud Run deployment on GCP (for cost efficiency and cold start performance)

Since this app gets infrequent traffic, serverless deployment makes perfect sense. I'm replacing the old Laravel <br>
admin panel with static codebases managed through version control rather than a database-driven admin interface.<br>
For that reason, we're adding the UI to manage photos directly to this codebase.

I use Docker for building the Go project but run it locally since Go runs on any architecture. For the database, <br>
I'm using SQLite with persistent storage mounted to the Cloud Run instance—GCP does not offer serverless database <br>
options, and performance isn't critical here.

Since Cloudflare doesn't provide local emulation for their Images service, I've built import/export functionality for <br>
the SQLite database and can purge unused images from Cloudflare. I'll start with a baseline containing only the <br>
current portfolio photos but retain the ability to share full-resolution shoots. This allows the flexibility of testing<br>
it locally using the real services, but easily prune out data that isn't necessary.

I might add Docker support later, but local tooling and deployment scripts work fine for now. This approach lets me <br>
use the Wrangler CLI without adding API keys to container environments—I can stick with local auth callbacks.

## Development

`cd frontend && pnpm install && pnpm run dev`

If you're using GoLand (or another Go IDE), just run the debug config against `main.go`. It'll prompt to make sure you<br>
have Go, and the xcode command line tools installed.

## Deployment

The backend can be built with Docker:
`docker build -f Dockerfile --target dum-e -o ./bin .`