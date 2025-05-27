import React, { useEffect, useState } from "react";
import silverTrophy from "../assets/trophy/silver.png";
import bronzeTrophy from "../assets/trophy/bronze.png";
import platinumTrophy from "../assets/trophy/platinum.png";
import goldTrophy from "../assets/trophy/gold.png";

interface User {
  id: number;
  name: string;
  email: string;
}

interface AchievementResponse {
  user: User;
  achievementLevel: string;
}

interface AchievementModalProps {
  userId: number;
  onClose: () => void;
}

const trophyMap: Record<string, string> = {
  Bronze: bronzeTrophy,
  Silver: silverTrophy,
  Gold: goldTrophy,
  Platinum: platinumTrophy,
};

const AchievementModal: React.FC<AchievementModalProps> = ({
  userId,
  onClose,
}) => {
  const [data, setData] = useState<AchievementResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    setLoading(true);
    setError(null);

    fetch(`http://localhost:5000/users/${userId}/achievement-level`)
      .then((res) => {
        if (!res.ok) throw new Error("Failed to fetch achievement data");
        return res.json();
      })
      .then((json) => {
        setData(json);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message || "Unknown error");
        setLoading(false);
      });
  }, [userId]);

  const trophySrc = data?.achievementLevel
    ? trophyMap[data.achievementLevel]
    : null;

  return (
    <div
      className="modal-overlay"
      onClick={onClose}
      style={{
        position: "fixed",
        inset: 0,
        backgroundColor: "rgba(0,0,0,0.7)",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        zIndex: 1050,
      }}
    >
      <div
        className="modal-content bg-dark text-light rounded"
        style={{
          width: "100%",
          maxWidth: "600px",
          position: "relative",
          padding: "2rem",
        }}
        onClick={(e) => e.stopPropagation()}
      >
        <button
          onClick={onClose}
          style={{
            position: "absolute",
            top: "12px",
            right: "12px",
            background: "transparent",
            border: "none",
            color: "white",
            fontSize: "1.5rem",
            cursor: "pointer",
          }}
          aria-label="Close modal"
        >
          &times;
        </button>

        {loading && <p>Loading achievement data...</p>}
        {error && <p className="text-danger">Error: {error}</p>}

        {data && (
          <div
            className="d-flex"
            style={{
              alignItems: "center",
              gap: "1.5rem",
              minHeight: "150px",
            }}
          >
            <div className="flex-grow-1">
              <h4 className="mb-2">{data.user.name}'s Achievement Level</h4>
              <h3
                className="mb-0"
                style={{
                  color: "#4db8ff",
                  fontWeight: "bold",
                  letterSpacing: "0.5px",
                }}
              >
                {data.achievementLevel}
              </h3>
            </div>

            {trophySrc && (
              <img
                src={trophySrc}
                alt={`${data.achievementLevel} trophy`}
                style={{
                  height: "85%",
                  maxHeight: "160px",
                  objectFit: "contain",
                }}
              />
            )}
          </div>
        )}
      </div>
    </div>
  );
};

export default AchievementModal;
