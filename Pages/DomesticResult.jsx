import React from "react";
import Standing from "./Components/Standings";
import LongestRuns from "./Components/LongestRuns";
import GameResults from "./Components/GameResults";
import UpcomingGames from "./Components/UpcomingGames";
import TopPlayers from "./Components/TopPlayers";

export default function DomesticResult () {
    return (
        <div className="domestic-result">
            <LongestRuns></LongestRuns>
            <GameResults></GameResults>
            <UpcomingGames></UpcomingGames>
            <Standing></Standing>
            <TopPlayers></TopPlayers>
        </div>       
    );
};