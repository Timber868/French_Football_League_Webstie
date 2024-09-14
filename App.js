import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";

import HomePage from "./Pages/HomePage/HomePage";
import GameResults from "./Pages/GameResultsPage/GameResultsPage";
import LeagueStandings from "./Pages/LeagueStandingsPage/LeagueStandings";
import PlayerStatsPage from "./Pages/PlayerStatsPage/PlayerStats";
import TeamPerformancePage from "./Pages/TeamPerformancesPage/TeamPerformancePage";
import UpcomingGamesPage from "./Pages/UpcomingGamesPage/UpcomingGames";
import NavBar from "./Pages/Components/NavBar/NavBar";

import "./App.scss";

export default function App() {
  return (
    <Router>
      <div className="App">
        <NavBar/>
        <div className="content">
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/game-results" element={<GameResults />} />
            <Route path="/standings" element={<LeagueStandings />} />
            <Route path="/player-stats" element={<PlayerStatsPage />} />
            <Route path="/team-performance" element={<TeamPerformancePage />} />
            <Route path="/upcoming-games" element={<UpcomingGamesPage />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
}
