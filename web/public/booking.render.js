let seats = [
    {"row": "A", "column": "1", "status": "available"},
    {"row": "A", "column": "2", "status": "occupied"},
    {"row": "A", "column": "3", "status": "available"},
    {"row": "A", "column": "4", "status": "occupied"},
    {"row": "A", "column": "5", "status": "occupied"},
    {"row": "A", "column": "6", "status": "available"},
    {"row": "B", "column": "1", "status": "available"},
    {"row": "B", "column": "2", "status": "occupied"},
    {"row": "B", "column": "3", "status": "available"},
    {"row": "B", "column": "4", "status": "occupied"},
    {"row": "B", "column": "5", "status": "occupied"},
    {"row": "B", "column": "6", "status": "available"},
    {"row": "C", "column": "1", "status": "occupied"},
    {"row": "C", "column": "2", "status": "available"},
    {"row": "C", "column": "3", "status": "occupied"},
    {"row": "C", "column": "4", "status": "available"},
    {"row": "C", "column": "5", "status": "available"},
    {"row": "C", "column": "6", "status": "occupied"},
    {"row": "D", "column": "1", "status": "available"},
    {"row": "D", "column": "2", "status": "occupied"},
    {"row": "D", "column": "3", "status": "available"},
    {"row": "D", "column": "4", "status": "occupied"},
    {"row": "D", "column": "5", "status": "occupied"},
    {"row": "D", "column": "6", "status": "available"},
    {"row": "E", "column": "1", "status": "available"},
    {"row": "E", "column": "2", "status": "occupied"},
    {"row": "E", "column": "3", "status": "available"},
    {"row": "E", "column": "4", "status": "occupied"},
    {"row": "E", "column": "5", "status": "available"},
    {"row": "E", "column": "6", "status": "occupied"},
    {"row": "F", "column": "1", "status": "available"},
    {"row": "F", "column": "2", "status": "occupied"},
    {"row": "F", "column": "3", "status": "available"},
    {"row": "F", "column": "4", "status": "occupied"},
    {"row": "F", "column": "5", "status": "occupied"},
    {"row": "F", "column": "6", "status": "available"},
    {"row": "G", "column": "1", "status": "occupied"},
    {"row": "G", "column": "2", "status": "available"},
    {"row": "G", "column": "3", "status": "occupied"},
    {"row": "G", "column": "4", "status": "available"},
    {"row": "G", "column": "5", "status": "occupied"},
    {"row": "G", "column": "6", "status": "available"},
    {"row": "G", "column": "7", "status": "occupied"},
    {"row": "G", "column": "8", "status": "available"},
    {"row": "G", "column": "9", "status": "occupied"},
    {"row": "G", "column": "10", "status": "available"}
];

// let seats = [
//     {"row": "A", "column": "1", "status": "available"},
//     {"row": "A", "column": "2", "status": "occupied"},
//     {"row": "A", "column": "3", "status": "available"},
//     {"row": "A", "column": "4", "status": "occupied"},
//     {"row": "A", "column": "5", "status": "occupied"},
//     {"row": "A", "column": "6", "status": "available"},
//     {"row": "B", "column": "1", "status": "available"},
//     {"row": "B", "column": "2", "status": "occupied"},
//     {"row": "B", "column": "3", "status": "available"},
//     {"row": "B", "column": "4", "status": "occupied"},
//     {"row": "B", "column": "5", "status": "occupied"},
//     {"row": "B", "column": "6", "status": "available"},
//     {"row": "C", "column": "1", "status": "occupied"},
//     {"row": "C", "column": "2", "status": "available"},
//     {"row": "C", "column": "3", "status": "occupied"},
//     {"row": "C", "column": "4", "status": "available"},
//     {"row": "C", "column": "5", "status": "available"},
//     {"row": "C", "column": "6", "status": "occupied"},
//     {"row": "D", "column": "1", "status": "available"},
//     {"row": "D", "column": "2", "status": "occupied"},
//     {"row": "D", "column": "3", "status": "available"},
//     {"row": "D", "column": "4", "status": "occupied"},
//     {"row": "D", "column": "5", "status": "occupied"},
//     {"row": "D", "column": "6", "status": "available"},
//     {"row": "D", "column": "7", "status": "occupied"},
//     {"row": "D", "column": "8", "status": "available"},
//     {"row": "D", "column": "9", "status": "occupied"},
//     {"row": "D", "column": "10", "status": "occupied"},
//     {"row": "D", "column": "11", "status": "available"},
// ];

document.addEventListener("DOMContentLoaded", () => {
    renderRows();
    renderSeats();
});


function renderRows() {
    const container = document.querySelector(".seats");

    if (container) {
        let rendered = [];

        seats.forEach(seat => {
            if (rendered.includes(seat.row)) {
                return;
            } else {
                rendered.push(seat.row);
    
                const rowDiv = document.createElement('div');
                rowDiv.classList.add('row');
                rowDiv.setAttribute('row', seat.row);
                container.appendChild(rowDiv);
            } 
        });
    }
}

function renderSeats() {
    seats.forEach(seat => {
        const rowContainer = document.querySelector(`div[row="${seat.row}"]`);

        if (rowContainer) {
                rowContainer.insertAdjacentHTML('beforeend', `
                    <div class="dropdown seat common ${seat.status}">
                        <a href="#" class="seat-button" role="button" data-bs-toggle="dropdown" aria-expanded="false"></a>
                        <ul class="dropdown-menu">
                            <li><a class="dropdown-item" href="#">Adult  -  1900тг</a></li>
                            <li><a class="dropdown-item" href="#">Kids  -  1450тг</a></li>
                            <li><a class="dropdown-item" href="#">Senior  -  1500тг</a></li>
                            <li><a class="dropdown-item" href="#">Student  -  1400тг</a></li>
                        </ul>
                    </div>
                `);
        }
    });
}