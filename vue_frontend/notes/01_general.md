 foxcodenine@foxcodenine-NUC12WSHi7:chess.log$ git add -A
 foxcodenine@foxcodenine-NUC12WSHi7:chess.log$ git commit -m 'init frontend'
[main 54f003b] init frontend
 48 files changed, 7376 insertions(+), 23 deletions(-)
 delete mode 100644 chess.log/src/App.vue
 delete mode 100644 chess.log/src/main.ts
 rename {chess.log => vue_frontend}/.gitignore (100%)
 rename {chess.log => vue_frontend}/.vscode/extensions.json (100%)
 rename {chess.log => vue_frontend}/README.md (100%)
 rename {chess.log => vue_frontend}/env.d.ts (100%)
 rename {chess.log => vue_frontend}/index.html (100%)
 create mode 100644 vue_frontend/notes/02_setup.md
 create mode 100644 vue_frontend/notes/03_node_vite_and_npm.md
 create mode 100644 vue_frontend/package-lock.json
 rename {chess.log => vue_frontend}/package.json (93%)
 rename {chess.log => vue_frontend}/public/favicon.ico (100%)
 create mode 100644 vue_frontend/src/App.vue
 create mode 100644 vue_frontend/src/assets/main.css
 create mode 100644 vue_frontend/src/components/TheSidebar.vue
 create mode 100644 vue_frontend/src/main.ts
 rename {chess.log => vue_frontend}/src/router/index.ts (100%)
 rename {chess.log => vue_frontend}/src/stores/counter.ts (100%)
 create mode 100755 vue_frontend/src/stores/dashboardStore.ts
 create mode 100755 vue_frontend/src/template.vue
 create mode 100755 vue_frontend/src/ui/components/_buttons.scss
 create mode 100755 vue_frontend/src/ui/components/_form.scss
 create mode 100755 vue_frontend/src/ui/components/_headers.scss
 create mode 100755 vue_frontend/src/ui/components/index.scss
 create mode 100755 vue_frontend/src/ui/index.scss
 create mode 100755 vue_frontend/src/ui/index.ts
 create mode 100755 vue_frontend/src/ui/primitives/VIconButton.vue
 create mode 100755 vue_frontend/src/ui/primitives/VModal.vue
 create mode 100755 vue_frontend/src/ui/primitives/VPager.vue
 create mode 100755 vue_frontend/src/ui/primitives/VSearch.vue
 create mode 100755 vue_frontend/src/ui/primitives/VTable.vue
 create mode 100755 vue_frontend/src/ui/primitives/VTable.vue.bak
 create mode 100755 vue_frontend/src/ui/primitives/VTabs.vue
 create mode 100755 vue_frontend/src/ui/primitives/VTabs_JS.vue
 create mode 100755 vue_frontend/src/ui/primitives/Vview.vue
 create mode 100755 vue_frontend/src/ui/svg/sprite.svg
 create mode 100755 vue_frontend/src/ui/tokens/_colors.scss
 create mode 100755 vue_frontend/src/ui/tokens/_extend.scss
 create mode 100755 vue_frontend/src/ui/tokens/_index.scss
 create mode 100755 vue_frontend/src/ui/tokens/_mixins.scss
 create mode 100755 vue_frontend/src/ui/tokens/_tailwind.scss
 create mode 100755 vue_frontend/src/ui/tokens/_typography.scss
 create mode 100755 vue_frontend/src/ui/tokens/_variables.scss
 create mode 100755 vue_frontend/src/ui/tokens/tailwind/_cursor.scss
 rename {chess.log => vue_frontend}/tsconfig.app.json (100%)
 rename {chess.log => vue_frontend}/tsconfig.json (100%)
 rename {chess.log => vue_frontend}/tsconfig.node.json (100%)
 rename {chess.log => vue_frontend}/vite.config.ts (72%)
 foxcodenine@foxcodenine-NUC12WSHi7:chess.log$ git-add-chrisfarrugiadev 
Enter passphrase for /home/foxcodenine/.ssh/chris_farrugia_dev_git/chris_farrugia_dev_git: 
Identity added: /home/foxcodenine/.ssh/chris_farrugia_dev_git/chris_farrugia_dev_git (chrisfarrugia.dev@gmail.com)
 foxcodenine@foxcodenine-NUC12WSHi7:chess.log$ git push
remote: Permission to ChrisFarrugiaDev/chess.log.git denied to foxcodenine.
fatal: unable to access 'https://github.com/ChrisFarrugiaDev/chess.log.git/': The requested URL returned error: 403
 foxcodenine@foxcodenine-NUC12WSHi7:chess.log$ git remote set-url origin git@github.com:ChrisFarrugiaDev/chess.log.git
 foxcodenine@foxcodenine-NUC12WSHi7:chess.log$ git remote -v
origin  git@github.com:ChrisFarrugiaDev/chess.log.git (fetch)
origin  git@github.com:ChrisFarrugiaDev/chess.log.git (push)
 foxcodenine@foxcodenine-NUC12WSHi7:chess.log$ git push
sign_and_send_pubkey: signing failed for ED25519 "/home/foxcodenine/.ssh/id_ed25519" from agent: agent refused operation
Enumerating objects: 75, done.
Counting objects: 100% (75/75), done.
Delta compression using up to 22 threads
Compressing objects: 100% (62/62), done.
Writing objects: 100% (75/75), 67.61 KiB | 11.27 MiB/s, done.
Total 75 (delta 5), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (5/5), done.
To github.com:ChrisFarrugiaDev/chess.log.git
 * [new branch]      main -> main