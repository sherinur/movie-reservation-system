let movies = [
    {
      "title": "Parthenope",
      "genre": "drama • fantasy",
      "description": "",
      "posterimage": "https://cdn.kino.kz/movies/Parthenope/p168x242.webp",
      "duration": 0,
      "language": "",
      "releasedate": "",
      "rating": "6.2",
      "pgrating": "16+",
      "production": "",
      "producer": "",
      "status": ""
    },
    {
      "title": "Babygirl",
      "genre": "drama • thriller",
      "description": "",
      "posterimage": "https://cdn.kino.kz/movies/Babygirl/p168x242.webp",
      "duration": 0,
      "language": "",
      "releasedate": "",
      "rating": "5.7",
      "pgrating": "18+",
      "production": "",
      "producer": "",
      "status": ""
    },
    {
      "title": "Escape",
      "genre": "thriller",
      "description": "",
      "posterimage": "https://cdn.kino.kz/movies/Escape/p168x242.webp",
      "duration": 0,
      "language": "",
      "releasedate": "",
      "rating": "3.9",
      "pgrating": "18+",
      "production": "",
      "producer": "",
      "status": ""
    },
    {
      "title": "Pemukiman Setan",
      "genre": "action movie • horror",
      "description": "",
      "posterimage": "https://cdn.kino.kz/movies/Pemukiman_Setan/p168x242.webp",
      "duration": 0,
      "language": "",
      "releasedate": "",
      "rating": "3.9",
      "pgrating": "18+",
      "production": "",
      "producer": "",
      "status": ""
    },
    {
      "title": "Alarum",
      "genre": "action movie • crime",
      "description": "",
      "posterimage": "https://cdn.kino.kz/movies/Alarum/p168x242.webp",
      "duration": 0,
      "language": "",
      "releasedate": "",
      "rating": "3.8",
      "pgrating": "18+",
      "production": "",
      "producer": "",
      "status": ""
    },
    {
      "title": "Better Man",
      "genre": "biography • drama",
      "description": "",
      "posterimage": "https://cdn.kino.kz/movies/Better_Man/p168x242.webp",
      "duration": 0,
      "language": "",
      "releasedate": "",
      "rating": "8.6",
      "pgrating": "18+",
      "production": "",
      "producer": "",
      "status": ""
    },
    {
      "title": "Den of Thieves 2: Pantera",
      "genre": "action movie • drama",
      "description": "",
      "posterimage": "https://cdn.kino.kz/movies/Den_of_Thieves_2__Pantera/p168x242.webp",
      "duration": 0,
      "language": "",
      "releasedate": "",
      "rating": "7.9",
      "pgrating": "16+",
      "production": "",
      "producer": "",
      "status": ""
    },
  ]

function displayMovies(movies) {
    const movieContainer = document.getElementById('movie-container');
    movies.forEach(movie => {
        const card = document.createElement('div');
        card.className = 'col-md-2 movie-card';
        card.innerHTML = `
            <a href="/reserve" style="text-decoration: none; color: inherit;" >
            <div class="card h-60>
                <div class="poster-container">
                    <img src="${movie.posterimage}" class="card-img-top" alt="${movie.title}">
                    <div class="rating-badge">
                        <span>&#9733;</span> ${movie.rating}
                    </div>
                    <div class="pg-rating">
                        ${movie.pgrating}
                    </div>
                </div>
                <div class="card-body">
                    <h6 class="card-title">${movie.title}</h6>
                    <p>${movie.genre}</p>
                </div>
            </div>
            <a>
        `;
        movieContainer.appendChild(card);
    });
}

displayMovies(movies);