#  Screenshot Search

> **"Your screenshots are your second memory."**

You've screenshotted 3,247 things. A dress, a restaurant, a tutorial, a boarding pass, a meme you'll never send. You will never find any of them again by scrolling.

**Screenshot Search** turns your camera roll's most neglected folder into a searchable memory layer — ask for it in plain English, get the screenshot back.

```
"the black dress under ₹2,000"        →  📱 ZARA Midi Dress, ₹1,499
"that restaurant I saved for Mumbai"  →  📱 Instagram post, Bandra
"the React tutorial I screenshotted"  →  📱 Blog post, saved 3 weeks ago
```

No folders. No manual tagging. No remembering the exact words in the image. Just ask.

---

## How it actually works

<table>
<tr><td>

**1. You screenshot, like always**
Nothing changes about your behavior. Keep taking screenshots exactly like today.

</td><td>

**2. We index it — then forget the image**
The screenshot is uploaded to a *temporary* buffer, read by AI (OCR + vision + embeddings), and immediately deleted. Only the *meaning* is kept — never the picture.

</td></tr>
<tr><td>

**3. You ask, in your own words**
"black dress I saved" works just as well as "ZARA midi dress ₹1499" — semantic search, not keyword matching.

</td><td>

**4. We hand you back the original**
Search resolves to your phone's own copy via its device asset ID. The screenshot never had a permanent home anywhere but your phone.

</td></tr>
</table>

<details>
<summary> <b>Click to see exactly what happens to your screenshot, step by step</b></summary>

```
  Your phone                                         Our backend
──────────────────                                    ─────────────────

 IMG_4821.png
      │
      │  1. upload (temporary, encrypted in transit)
      ▼
                                                 ┌───────────────────┐
                                                 │  Temporary Storage │
                                                 │   (processing only)│
                                                 └─────────┬─────────┘
                                                           │  2. worker picks it up
                                                           ▼
                                                  ┌─────────────────┐
                                                  │  Gemini Vision    │
                                                  │  "black midi      │
                                                  │   dress, ZARA,    │
                                                  │   ₹1,499"         │
                                                  └────────┬─────────┘
                                                           │  3. + embedding
                                                           ▼
                                                  ┌─────────────────┐
                                                  │   PostgreSQL      │
                                                  │   + pgvector       │
                                                  │                    │
                                                  │  ✅ description    │
                                                  │  ✅ OCR text       │
                                                  │  ✅ embedding      │
                                                  │  ✅ device_asset_id│
                                                  │  ❌ the image      │ ← deleted, always
                                                  └─────────────────┘
      │
      │  original stays right here, forever
      ▼
 IMG_4821.png  (untouched)
```

**The rule we don't break:** if it's not derived metadata, it doesn't survive indexing. The image is a guest, not a tenant.

</details>

---

##  Where we are

```
  ██████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░  Phase 1 of 5 — scaffold complete
```

| Phase | What it proves | Status |
|---|---|---|
| **1 — Backend Foundation** | Go API + Postgres + auth actually run | 🔨 scaffolded, building now |
| **2 — AI Indexing** | A real screenshot becomes searchable metadata | ⏳ next |
| **3 — Search** | A sentence finds the right screenshot | ⏳ |
| **4 — Mobile** | It works on an actual phone | ⏳ |
| **5 — Polish** | It's demo-ready and privacy-honest | ⏳ |

We build **one version at a time** — V1 is the 20–100-screenshot proof that the whole loop works, not the finished product. V2 (categories, dedup, collections) and V3 (personal-memory queries like *"what did I save about system design last month?"*) come later, deliberately.

---

##  The map, if you're poking around the code

```
screenshot-search/
├── cmd/server/main.go          ← boots the API (health check lives here today)
├── internal/
│   ├── handler/                ← HTTP layer only — no business logic allowed in here
│   │   ├── auth.go
│   │   ├── screenshot.go
│   │   └── search.go
│   ├── service/                ← the brains: AI orchestration, search, sync logic
│   │   ├── ai_service.go       ← the ONE interface Gemini hides behind
│   │   ├── screenshot_service.go
│   │   └── search_service.go
│   ├── worker/                 ← async indexing pool — the API never blocks on AI
│   │   └── indexing_worker.go
│   ├── repository/             ← the only code allowed to speak SQL
│   │   └── screenshot_repository.go
│   ├── storage/                ← temporary object storage, never permanent
│   │   └── storage.go
│   ├── model/                  ← what a Screenshot *is*
│   │   └── screenshot.go
│   └── config/                 ← env vars in, typed config out
├── migrations/                 ← schema, versioned
├── Dockerfile                  ← written today, not used until the laptop can run it
└── .env.example                ← copy to .env, fill in secrets, never commit it
```

**One rule that shapes everything above:** the mobile app never hears the word "Gemini." It sees `upload → indexed ✓`. Every AI decision — which provider, which prompt, which embedding strategy — lives behind `AIProvider` in `ai_service.go`. Swap Gemini for anything else later and nothing outside that file notices.

---

##  Running it locally

```bash
cp .env.example .env        # fill in DATABASE_URL, JWT_SECRET, GEMINI_API_KEY later
go build ./...               # confirms everything compiles
go run ./cmd/server           # starts on :8080
curl localhost:8080/health    # → {"status":"ok"}
```

No Docker required — and none is required *by design* until this machine can run it. The architecture stays container-ready (see `Dockerfile`), but every phase up through mobile integration runs as a native Go binary against a local Postgres.

---

##  The three decisions worth knowing about

<table>
<tr>
<td width="33%" valign="top">

###  Screenshots aren't stored
They're processed and forgotten. Object storage is a *buffer*, not a bucket. This isn't a compliance checkbox — screenshots routinely contain OTPs, bank screenshots, private chats. We designed around that from day one, not as an afterthought.

</td>
<td width="33%" valign="top">

###  The AI provider is swappable
`AIProvider` is one small interface: analyze an image, embed text, embed an image. Gemini implements it today. Nothing else in the codebase is allowed to know that.

</td>
<td width="33%" valign="top">

###  Identity is a content hash, not a filename
Rename `IMG_4821.png` to `favorite-dress.png` and we still know it's the same screenshot — `content_hash`, not `device_asset_id` alone, is what makes incremental indexing ("3,247 → 3,248 processes *only* the new one") actually correct.

</td>
</tr>
</table>

---

##  Why this exists

Somewhere in your camera roll right now is a dress you wanted, a place you meant to visit, and a piece of advice you screenshotted and never read again.

This project is the bet that **search**, not **folders**, is how you'll find it.
