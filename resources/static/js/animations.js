var $ = (content) => document.querySelector(content);
var $$ = (content) => document.querySelectorAll(content);
var reveals = $$('[data-fade-in]');
var fadeLefts = $$('[data-fade-left]');
var fadeRights = $$('[data-fade-right]');
var fadeUps = $$('[data-fade-up]');
var counts = $$('[data-count]');
var observerContructor = (animation) => {
    return new IntersectionObserver((entries, observer) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.classList.add(animation);
                observer.unobserve(entry.target);
            }
        });
    })
}
var revealObserver = observerContructor('fade-in');
var fadeLeftObserver = observerContructor('fade-left');
var fadeRightObserver = observerContructor('fade-right');
var fadeUpObserver = observerContructor('fade-up');
var counterObserver = new IntersectionObserver((entries, observer) => {
    entries.forEach(entry => {
        if (entry.isIntersecting) {
            const num = parseInt(entry.target.getAttribute('data-count'));
            const baseNum = num / 20
            let i = 1
            let intervalId = null
            intervalId = setInterval(() => {
                entry.target.innerHTML = Math.floor(baseNum * i);
                i++
                if (i > 20) {
                    clearInterval(intervalId);
                }
            }, 50)
            observer.unobserve(entry.target);
        }
    })
})

reveals.forEach(reveal => {
    revealObserver.observe(reveal);
});
fadeLefts.forEach(reveal => {
    fadeLeftObserver.observe(reveal);
});
fadeRights.forEach(reveal => {
    fadeRightObserver.observe(reveal);
});
fadeUps.forEach(reveal => {
    fadeUpObserver.observe(reveal);
});

counts.forEach(reveal => {
    counterObserver.observe(reveal);
});

