import DomesticSection from "../Components/BuildingBlocs/DomesticSections";

export default function HomePage() {
    return (
    <div className="home-page">
         <div>
      {/* Introduction */}
      <header>
        <h1>Welcome to the Ligue 1 Website</h1>
        <p>Your go-to source for Ligue 1 standings, top players, and more!</p>
      </header>

      {/* Quick Stats Overview */}
      <section>
        <h2>League Standings (Top 5)</h2>
        <ul>
          <li>1. PSG - 45 points</li>
          <li>2. Olympique Lyonnais - 40 points</li>
          <li>3. AS Monaco - 39 points</li>
          <li>4. Lille OSC - 36 points</li>
          <li>5. Marseille - 34 points</li>
        </ul>
      </section>

      <section>
        <h2>Top Scorers</h2>
        <ul>
          <li>Kylian Mbappé - 20 goals</li>
          <li>Wissam Ben Yedder - 18 goals</li>
          <li>Moussa Dembélé - 15 goals</li>
        </ul>
      </section>

      <section>
        <h2>Upcoming League Games</h2>
        <ul>
          <li>PSG vs Olympique Lyonnais - 2024-09-15</li>
          <li>AS Monaco vs Marseille - 2024-09-16</li>
          <li>Lille OSC vs Montpellier - 2024-09-17</li>
        </ul>
      </section>

      <section>
        <h2>Upcoming European Games</h2>
        <ul>
          <li>PSG vs Olympique Lyonnais - 2024-09-15</li>
          <li>AS Monaco vs Marseille - 2024-09-16</li>
          <li>Lille OSC vs Montpellier - 2024-09-17</li>
        </ul>
      </section>

    </div>
    </div>
    );
}