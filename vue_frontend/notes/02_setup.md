 
### Step 1. Create a Vue Project with vite

 ```bash
 foxcodenine@foxcodenine-NUC12WSHi7:chess.log$ npm init vite@latest

> npx
> create-vite

│
◇  Project name:
│  chess.log
│
◇  Select a framework:
│  Vue
│
◇  Select a variant:
│  Official Vue Starter ↗
│
◇  Use rolldown-vite (Experimental)?:
│  No
│
◇  Install with npm and start now?
│  Yes
Need to install the following packages:
create-vue@3.18.2
Ok to proceed? (y) 


> npx
> create-vue chess.log

┌  Vue.js - The Progressive JavaScript Framework
│
◇  Select features to include in your project: (↑/↓ to navigate, space to select, a to toggle all, enter to confirm)
│  TypeScript, Router (SPA development), Pinia (state management)
│
◇  Select experimental features to include in your project: (↑/↓ to navigate, space to select, a to toggle all, enter to confirm)
│  none
│
◇  Skip all example code and start with a blank Vue project?
│  Yes

Scaffolding project in /home/foxcodenine/foxfiles/git/chrisfarrugia.dev/chess.log/chess.log...
│
└  Done. Now run:

   cd chess.log
   npm install
   npm run dev

| Optional: Initialize Git in your project directory with:
  
   git init && git add -A && git commit -m "initial commit"


```

### Step 2. Install sass:

```bash
npm install sass@latest sass-loader@latest
npm install -D sass-embedded
```