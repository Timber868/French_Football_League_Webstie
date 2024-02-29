import DomesticSection from "./BuildingBlocs/DomesticSections";
import LongestLosing from "./Streaks/LongestLosing";
import LongestUnbeaten from "./Streaks/LongestUnbeaten";
import LongestWinning from "./Streaks/LongestWinning";

export default function LongestRuns () {
    return(
        <DomesticSection title={"Longest Runs"} className="longest-runs">
            <LongestLosing></LongestLosing>
            <LongestUnbeaten></LongestUnbeaten>
            <LongestWinning></LongestWinning>
        </DomesticSection>
    );
};