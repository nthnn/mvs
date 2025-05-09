/*
 * Copyright 2025 Nathanne Isip
 * This file is part of Minimal Versioning System (https://github.com/nthnn/mvs)
 * This code is licensed under MIT license (see LICENSE for details)
 */

const contentEl = document.getElementById('content');
const progressEl = document.getElementById('progress');

function simulateProgress(duration) {
    let start = Date.now();

    progressEl.max = 100;
    progressEl.value = 0;
    progressEl.style.display = 'block';
  
    return setInterval(() => {
        const elapsed = Date.now() - start;
        const percent = Math.min(
            (elapsed / duration) * 100,
            100
        );

        progressEl.value = percent;
    }, 50);
}

(async()=> {
    const duration = 2000;
    const animId = simulateProgress(duration);

    const fetchPromise = fetch(
        "https://raw.githubusercontent.com/nthnn/mvs/refs/heads/main/README.md"
    ).then(res => {
        if(!res.ok)
            throw new Error(`Cannot fetch content data`);

        return res.text();
    });

    try {
        const [markdown] = await Promise.all([
            fetchPromise,
            new Promise(r => setTimeout(r, duration))
        ]);

        contentEl.innerHTML = marked.parse(markdown);
        contentEl.querySelectorAll('a').forEach(a =>
            a.setAttribute('target', '_blank')
        );
    }
    catch(_) {
        contentEl.innerHTML = `<br/><br/><p align="center"><img src="https://raw.githubusercontent.com/nthnn/mvs/refs/heads/main/assets/mvs-logo.png" width="250" /><br/><br/><span style="color:red;">Error:</span> ${error.message}</p>`;
    }
    finally {
        clearInterval(animId);
        progressEl.style.display = 'none';
    }
})();
