import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import './NavBar.scss';

export default function NavBar() {
    return (
        <nav className="navbar">
            <div id="Ligue1Logo"> Ligue 1 </div>
            <ul className="page-links">
                <li><Link to="/" className="page-button">Home</Link></li>
                <li><Link to="/game-results" className="page-button">Game Results</Link></li>
                <li><Link to="/standings" className="page-button">League Standings</Link></li>
                <li><Link to="/player-stats" className="page-button">Player Stats</Link></li>
                <li><Link to="/team-performance" className="page-button">Team Performances</Link></li>
                <li><Link to="/upcoming-games" className="page-button">Upcoming Games</Link></li>
            </ul>
        </nav>
    );
}