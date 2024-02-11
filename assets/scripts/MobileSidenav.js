const sidenav = document.getElementById('sidenav');
const sidenavOpenBtn = document.getElementById('sidenav-open-btn');
const sidenavCloseBtn = document.getElementById('sidenav-close-btn');

sidenavOpenBtn.addEventListener("click", () => {
    sidenav.dataset.open = "true";
});

sidenavCloseBtn.addEventListener("click", () => {
    sidenav.dataset.open = "false";
});
