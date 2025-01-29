htmx.onLoad(function (content) {
    var $ = (c) => content.querySelector(c);
    var $$ = (c) => content.querySelectorAll(c);
    var reveals = $$('[data-fade-in]');
    var fadeLefts = $$('[data-fade-left]');
    var fadeRights = $$('[data-fade-right]');
    var fadeUps = $$('[data-fade-up]');
    var counts = $$('[data-count]');
    var modalTriggers = $$('[data-modal-trigger]')
    console.log(modalTriggers)

    modalTriggers.forEach(trigger => {
        trigger.addEventListener('click', function() {
            const modalId = this.getAttribute('data-modal-trigger');
            const modal = document.querySelector(`[data-modal="${modalId}"]`);
            if (modal) {
                modal.hidden = false; // Remove hidden property
            }
        });
    });


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
})
