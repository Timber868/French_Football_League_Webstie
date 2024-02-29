import DomesticSection from "./BuildingBlocs/DomesticSections";
import TopAssisters from "./TopPlayers/TopAssisters";
import TopScorers from "./TopPlayers/TopScorers";

export default function TopPlayers () {
    return (
        <DomesticSection title={"Top Player Stats"} className="top-players">
            <TopScorers></TopScorers>
            <TopAssisters></TopAssisters>
        </DomesticSection>
    );
};