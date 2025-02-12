import React, { useEffect, useState } from 'react';
import logo from '../src/logo.png';
import 'bootstrap/dist/css/bootstrap.min.css';
import './style.scss';

const MoviePage = () => {
  // const [movies, setMovies] = useState([]);
  const [loading, setLoading] = useState(true);

  // useEffect(() => {
  //   fetch("http://localhost/movi/movie")
  //     .then(response => response.json())
  //     .then(data => {
  //       setMovies(data); 
  //       setLoading(false);
  //     })
  //     .catch(error => {
  //       console.error("Movie download error:", error);
  //       setLoading(false);
  //     });
  // }, []);
let movies = [
  {
      "Title": "The Shawshank Redemption",
      "Genre": "Drama",
      "Description": "Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BMDAyY2FhYjctNDc5OS00MDNlLThiMGUtY2UxYWVkNGY2ZjljXkEyXkFqcGc@._V1_QL75_UX140_CR0,1,140,207_.jpg",
      "Duration": 142,
      "Language": "English",
      "ReleaseDate": "1994",
      "Rating": "9.3",
      "PGrating": "R",
      "Production": "Castle Rock Entertainment",
      "Producer": "Niki Marvin",
      "Status": "Released"
    
  },
  {
      "Title": "The Godfather",
      "Genre": "Crime, Drama",
      "Description": "The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BNGEwYjgwOGQtYjg5ZS00Njc1LTk2ZGEtM2QwZWQ2NjdhZTE5XkEyXkFqcGc@._V1_QL75_UY207_CR3,0,140,207_.jpg",
      "Duration": 175,
      "Language": "English, Italian",
      "ReleaseDate": "1972",
      "Rating": "9.2",
      "PGrating": "R",
      "Production": "Paramount Pictures",
      "Producer": "Albert S. Ruddy",
      "Status": "Released"
  },
  {
      "Title": "The Chaos Class",
      "Genre": "Comedy",
      "Description": "Lazy, uneducated students share a very close bond. They live together in the dormitory, where they plan their latest pranks. When a new headmaster arrives, the students naturally try to overthrow him. A comic war of nitwits follows.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BZTdhN2ViYjctMGZlZi00ZGRmLWIxMWQtNzJiMDY1ZDcxNjJmXkEyXkFqcGc@._V1_QL75_UX140_CR0,1,140,207_.jpg",
      "Duration": 85,
      "Language": "English, Italian",
      "ReleaseDate": "1975",
      "Rating": "9.2",
      "PGrating": "R",
      "Production": "Unknown",
      "Producer": "Ertem Egilmez",
      "Status": "Released"
  },
  {
      "Title": "The Dark Knight",
      "Genre": "Action, Crime, Drama",
      "Description": "When a menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman, James Gordon and Harvey Dent must work together to put an end to the madness.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BMTMxNTMwODM0NF5BMl5BanBnXkFtZTcwODAyMTk2Mw@@._V1_QL75_UX140_CR0,0,140,207_.jpg",
      "Duration": 152,
      "Language": "English",
      "ReleaseDate": "2008",
      "Rating": "9.0",
      "PGrating": "PG-13",
      "Production": "Warner Bros.",
      "Producer": "Christopher Nolan, Emma Thomas",
      "Status": "Released"
  },
  {
      "Title": "Schindler's List",
      "Genre": "Biography, Drama, History",
      "Description": "In German-occupied Poland during World War II, industrialist Oskar Schindler gradually becomes concerned for his Jewish workforce after witnessing their persecution by the Nazis.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BNjM1ZDQxYWUtMzQyZS00MTE1LWJmZGYtNGUyNTdlYjM3ZmVmXkEyXkFqcGc@._V1_QL75_UX140_CR0,1,140,207_.jpg",
      "Duration": 195,
      "Language": "English, German, Hebrew, Polish",
      "ReleaseDate": "1993",
      "Rating": "9.0",
      "PGrating": "R",
      "Production": "Universal Pictures",
      "Producer": "Steven Spielberg, Branko Lustig, Gerald R. Molen",
      "Status": "Released"
  },
  {
      "Title": "The Lord of the Rings: The Return of the King",
      "Genre": "Action, Adventure, Drama",
      "Description": "Gandalf and Aragorn lead the World of Men against Sauron's army to draw his gaze from Frodo and Sam as they approach Mount Doom with the One Ring.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BMTZkMjBjNWMtZGI5OC00MGU0LTk4ZTItODg2NWM3NTVmNWQ4XkEyXkFqcGc@._V1_QL75_UX140_CR0,0,140,207_.jpg",
      "Duration": 201,
      "Language": "English",
      "ReleaseDate": "2003",
      "Rating": "9.0",
      "PGrating": "PG-13",
      "Production": "New Line Cinema",
      "Producer": "Peter Jackson, Fran Walsh, Philippa Boyens",
      "Status": "Released"
  },
  {
      "Title": "12 Angry Men",
      "Genre": "Drama",
      "Description": "The jury in a New York City murder trial is frustrated by a single member whose skeptical caution forces them to more carefully consider the evidence before jumping to a hasty verdict.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BYjE4NzdmOTYtYjc5Yi00YzBiLWEzNDEtNTgxZGQ2MWVkN2NiXkEyXkFqcGc@._V1_QL75_UX140_CR0,4,140,207_.jpg",
      "Duration": 96,
      "Language": "English",
      "ReleaseDate": "1957",
      "Rating": "9.0",
      "PGrating": "R",
      "Production": "United Artists",
      "Producer": "Sidney Lumet, Reginald Rose",
      "Status": "Released"
  },
  {
      "Title": "12 Angry Men",
      "Genre": "Drama",
      "Description": "The jury in a New York City murder trial is frustrated by a single member whose skeptical caution forces them to more carefully consider the evidence before jumping to a hasty verdict.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BYjE4NzdmOTYtYjc5Yi00YzBiLWEzNDEtNTgxZGQ2MWVkN2NiXkEyXkFqcGc@._V1_QL75_UX140_CR0,4,140,207_.jpg",
      "Duration": 96,
      "Language": "English",
      "ReleaseDate": "1957",
      "Rating": "9.0",
      "PGrating": "R",
      "Production": "United Artists",
      "Producer": "Reginald Rose",
      "Status": "Released"
  },
  {
      "Title": "The Godfather Part II",
      "Genre": "Crime, Drama",
      "Description": "The early life and career of Vito Corleone in 1920s New York City is portrayed, while his son, Michael, expands and tightens his grip on the family crime syndicate.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BNzcwZWUzOWItMjcyYi00OWUwLThmZjUtY2M0ZjhjMzJiNmM4XkEyXkFqcGc@._V1_QL75_UX140_CR0,2,140,207_.jpg",
      "Duration": 202,
      "Language": "English, Italian",
      "ReleaseDate": "1974",
      "Rating": "9.0",
      "PGrating": "R",
      "Production": "Paramount Pictures",
      "Producer": "Francis Ford Coppola",
      "Status": "Released"
  },
  {
      "Title": "The Lord of the Rings: The Fellowship of the Ring",
      "Genre": "Adventure, Drama, Fantasy",
      "Description": "A meek Hobbit from the Shire and eight companions set out on a journey to destroy the powerful One Ring and save Middle-earth from the Dark Lord Sauron.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BNzIxMDQ2YTctNDY4MC00ZTRhLTk4ODQtMTVlOWY4NTdiYmMwXkEyXkFqcGc@._V1_QL75_UX140_CR0,0,140,207_.jpg",
      "Duration": 178,
      "Language": "English",
      "ReleaseDate": "2001",
      "Rating": "8.9",
      "PGrating": "PG-13",
      "Production": "New Line Cinema",
      "Producer": "Barrie M. Osborne, Peter Jackson, Fran Walsh",
      "Status": "Released"
  },
  {
      "Title": "Pulp Fiction",
      "Genre": "Crime, Drama",
      "Description": "The lives of two mob hitmen, a boxer, a gangster and his wife, and a pair of diner bandits intertwine in four tales of violence and redemption.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BYTViYTE3ZGQtNDBlMC00ZTAyLTkyODMtZGRiZDg0MjA2YThkXkEyXkFqcGc@._V1_QL75_UY207_CR1,0,140,207_.jpg",
      "Duration": 154,
      "Language": "English",
      "ReleaseDate": "1994",
      "Rating": "8.9",
      "PGrating": "R",
      "Production": "A Band Apart",
      "Producer": "Lawrence Bender",
      "Status": "Released"
  },
  {
      "Title": "Inception",
      "Genre": "Action, Adventure, Sci-Fi",
      "Description": "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a C.E.O., but his tragic past may doom the project and his team to disaster.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BMjAxMzY3NjcxNF5BMl5BanBnXkFtZTcwNTI5OTM0Mw@@._V1_QL75_UX140_CR0,0,140,207_.jpg",
      "Duration": 148,
      "Language": "English",
      "ReleaseDate": "2010",
      "Rating": "8.8",
      "PGrating": "PG-13",
      "Production": "Warner Bros. Pictures",
      "Producer": "Christopher Nolan, Emma Thomas",
      "Status": "Released"
  },
  {
      "Title": "Fight Club",
      "Genre": "Drama",
      "Description": "An insomniac office worker and a devil-may-care soap maker form an underground fight club that evolves into much more.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BOTgyOGQ1NDItNGU3Ny00MjU3LTg2YWEtNmEyYjBiMjI1Y2M5XkEyXkFqcGc@._V1_QL75_UX140_CR0,1,140,207_.jpg",
      "Duration": 139,
      "Language": "English",
      "ReleaseDate": "1999",
      "Rating": "8.8",
      "PGrating": "R",
      "Production": "20th Century Fox",
      "Producer": "Art Linson, Ceán Chaffin, Ross Grayson Bell",
      "Status": "Released"
  },
  {
      "Title": "Forrest Gump",
      "Genre": "Drama, Romance",
      "Description": "The history of the United States from the 1950s to the '70s unfolds from the perspective of an Alabama man with an IQ of 75, who yearns to be reunited with his childhood sweetheart.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BNDYwNzVjMTItZmU5YS00YjQ5LTljYjgtMjY2NDVmYWMyNWFmXkEyXkFqcGc@._V1_QL75_UY207_CR2,0,140,207_.jpg",
      "Duration": 142,
      "Language": "English",
      "ReleaseDate": "1994",
      "Rating": "8.8",
      "PGrating": "PG-13",
      "Production": "Paramount Pictures",
      "Producer": "Wendy Finerman, Steve Tisch, Steve Starkey",
      "Status": "Released"
  },
  {
      "Title": "The Lord of the Rings: The Two Towers",
      "Genre": "Action, Adventure, Drama",
      "Description": "While Frodo and Sam edge closer to Mordor with the help of the shifty Gollum, the divided fellowship makes a stand against Sauron's new ally, Saruman, and his hordes of Isengard.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BMGQxMDdiOWUtYjc1Ni00YzM1LWE2NjMtZTg3Y2JkMjEzMTJjXkEyXkFqcGc@._V1_QL75_UX140_CR0,5,140,207_.jpg",
      "Duration": 179,
      "Language": "English",
      "ReleaseDate": "2002",
      "Rating": "8.8",
      "PGrating": "PG-13",
      "Production": "New Line Cinema",
      "Producer": "Barrie M. Osborne, Fran Walsh, Peter Jackson",
      "Status": "Released"
  },
  {
      "Title": "Interstellar",
      "Genre": "Adventure, Drama, Sci-Fi",
      "Description": "When Earth becomes uninhabitable in the future, a farmer and ex-NASA pilot, Joseph Cooper, is tasked to pilot a spacecraft, along with a team of researchers, to find a new planet for humans.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BYzdjMDAxZGItMjI2My00ODA1LTlkNzItOWFjMDU5ZDJlYWY3XkEyXkFqcGc@._V1_QL75_UX140_CR0,0,140,207_.jpg",
      "Duration": 169,
      "Language": "English",
      "ReleaseDate": "2014",
      "Rating": "8.7",
      "PGrating": "PG-13",
      "Production": "Paramount Pictures, Warner Bros. Pictures, Legendary Pictures, Syncopy",
      "Producer": "Christopher Nolan, Emma Thomas, Lynda Obst",
      "Status": "Released"
  },
  {
      "Title": "The Green Mile",
      "Genre": "Crime, Drama, Fantasy",
      "Description": "Paul Edgecomb, the head death row guard at a prison in 1930s Louisiana, meets an inmate, John Coffey, a black man who is accused of murdering two girls. His life changes drastically when he discovers that John has a special gift.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BMTUxMzQyNjA5MF5BMl5BanBnXkFtZTYwOTU2NTY3._V1_QL75_UX140_CR0,0,140,207_.jpg",
      "Duration": 189,
      "Language": "English",
      "ReleaseDate": "1999",
      "Rating": "8.6",
      "PGrating": "R",
      "Production": "Castle Rock Entertainment",
      "Producer": "David Valdes, Frank Darabont",
      "Status": "Released"
  },
  {
      "Title": "Terminator 2: Judgment Day",
      "Genre": "Action, Sci-Fi",
      "Description": "A cyborg, identical to the one who failed to kill Sarah Connor, must now protect her ten year old son John from an even more advanced and powerful cyborg.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BNGMyMGNkMDUtMjc2Ni00NWFlLTgyODEtZTY2MzBiZTg0OWZiXkEyXkFqcGc@._V1_QL75_UX140_CR0,0,140,207_.jpg",
      "Duration": 137,
      "Language": "English",
      "ReleaseDate": "1991",
      "Rating": "8.6",
      "PGrating": "R",
      "Production": "Carolco Pictures",
      "Producer": "James Cameron, Gale Anne Hurd",
      "Status": "Released"
  },
  {
      "Title": "Gladiator",
      "Genre": "Action, Adventure, Drama",
      "Description": "A former Roman General sets out to exact vengeance against the corrupt emperor who murdered his family and sent him into slavery.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BYWQ4YmNjYjEtOWE1Zi00Y2U4LWI4NTAtMTU0MjkxNWQ1ZmJiXkEyXkFqcGc@._V1_QL75_UX140_CR0,0,140,207_.jpg",
      "Duration": 155,
      "Language": "English",
      "ReleaseDate": "2000",
      "Rating": "8.5",
      "PGrating": "R",
      "Production": "DreamWorks Pictures, Universal Pictures, Scott Free Productions",
      "Producer": "David Franzoni, Branko Lustig, Douglas Wick",
      "Status": "Released"
  },
  {
      "Title": "The Lion King",
      "Genre": "Animation, Adventure, Drama",
      "Description": "Lion prince Simba and his father are targeted by his bitter uncle, who wants to ascend the throne himself.",
      "PosterImage": "https://m.media-amazon.com/images/M/MV5BZGRiZDZhZjItM2M3ZC00Y2IyLTk3Y2MtMWY5YjliNDFkZTJlXkEyXkFqcGc@._V1_QL75_UX140_CR0,0,140,207_.jpg",
      "Duration": 88,
      "Language": "English",
      "ReleaseDate": "1994",
      "Rating": "8.5",
      "PGrating": "G",
      "Production": "Walt Disney Feature Animation",
      "Producer": "Don Hahn",
      "Status": "Released"
  }
]
  return (
    <div>
      <div className="background-wrapper"></div>

      <header className="text-white py-3">
        <div className="container d-flex justify-content-between align-items-center">
          <div className="d-flex align-items-center">
            <img src={logo} className="card-img-top" alt="logo" />
          </div>
          <div> 
            <a href="/login" className="btn btn-outline-light custom-green-btn">Log In</a>
            <a href="/register" className="btn btn-outline-light me-2">Register</a>
          </div>
        </div>
      </header>

      <div className="container text-center">
        <div className="row justify-content-center">
          <h2 className="mb-5 custom-white">Now Showing</h2>

          {/* {loading ? (
            <p className="text-white">Loading...</p>
          ) : movies.length === 0 ? (
            <p className="text-white">There are no available movies.</p>
          ) : ( */}
           {( movies.map((movie) => (
              <div className="col-md-2 movie-card" key={movie.id}>
                <a href="#" className="movie-link">
                  <div className="poster-container">
                    <img 
                      src={movie.PosterImage || "https://via.placeholder.com/140x207"} 
                      className="card-img-top" 
                      alt={movie.Title} 
                    />
                    <div className="rating-badge">
                      <span>&#9733;</span> {movie.Rating || "N/A"}
                    </div>
                    <div className="pg-rating">{movie.PGrating || "N/A"}+</div>
                  </div>
                  <div className="card-body">
                    <h6 className="card-title">{movie.Title}</h6>
                  </div>
                </a>
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  );
};

export default MoviePage;
