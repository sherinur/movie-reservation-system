document.addEventListener("DOMContentLoaded", () => {
    const dropdowns = document.querySelectorAll(".dropdown");

    dropdowns.forEach((dropdown) => {
        dropdown.addEventListener("click", (event) => {
            const clickedElement = event.target;
            if (clickedElement.classList.contains("dropdown-item")) {
                dropdown.classList.remove("common","love", "vip")
                dropdown.classList.add("selected");
            }
        });
    });

    dropdowns.forEach((dropdown) => {
        if (dropdown.classList.contains("occupied")) {
            const button = dropdown.querySelector("a")
            button.classList.remove("seat-button")
        } else {
            const button = dropdown.querySelector("a")
            button.classList.add("seat-button")
        }
    });
});
