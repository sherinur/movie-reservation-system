import express from 'express';
import 'dotenv/config';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const PORT = process.env.PORT || 4200;

const app = express();

async function main() {
    app.use(express.json());
    app.use(express.urlencoded({ extended: false }));
    app.use(express.static(path.join(__dirname, 'public')));

    app.get('/', (req, res) => {
        res.sendFile(path.join(__dirname, 'views', 'index.html'));
    });
    
    app.get('/reserve', (req, res) => {
        res.sendFile(path.join(__dirname, 'views', 'booking.html'))
    });
    
    app.get('/login', (req, res) => {
        res.sendFile(path.join(__dirname, 'views', 'login.html'))
    });

    app.get('/register', (req, res) => {
        res.sendFile(path.join(__dirname, 'views', 'register.html'))
    });

    app.listen(PORT, () => {
        console.log('Web server is running on port 4200');
    });
}

main()