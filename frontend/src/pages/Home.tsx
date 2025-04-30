import React, { useState, useEffect, useRef } from "react";
import CharacterCard from "../components/CharacterCard.tsx";
import "../styles/animations.css";
import "../styles/CharacterCard.css";
import { Character } from "../types/character.ts"
import SearchBar from "../components/SearchBar.tsx";
import CharacterDetails from "../components/CharacterDetails.tsx";
import useCharacterFetch from "../hooks/useFetchCharacter.tsx";


const Home: React.FC = () => {
    const [selectedChar, setSelectedChar] = useState<null | Character>(null);
    const containerRef = useRef<HTMLDivElement | null>(null);

    const {
        characters,
        filteredCharacters,
        loading,
        debouncedSearch,
        debouncedLoad,
        checkOverflow,
        handleScroll
    } = useCharacterFetch();

    const selectChar = (char: Character) => {
        setSelectedChar(char);
    };

    useEffect(() => {
        // Load initial characters
        debouncedLoad()
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [])

    useEffect(() => {
        // Initial data fetch
        handleScroll(containerRef);
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [characters]);

    useEffect(() => {
        checkOverflow(containerRef)
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [characters])

    return (
        <>
            <header>
                <h1>Welcome to the Marvel Universe!</h1>
                <p>Explore marvel characters and their comics</p>
            </header>
            <section id="charSection">
                <h2>Choose your character!</h2>
                <SearchBar onSearch={debouncedSearch} />
                <div className="charTab" ref={containerRef} onScroll={() => {
                    if (containerRef.current) {
                        handleScroll(containerRef);
                    }
                }}>
                    {(filteredCharacters.length > 0 ? filteredCharacters : characters).map((char) => (
                        <CharacterCard
                            key={char.id}
                            character={char}
                            onClick={selectChar}
                        />))
                    }
                    { /* Skeleton loading */
                        loading && Array(5).fill(0).map((_, index) => (
                            <div key={`skeleton-${index}`} className="character-card skeleton">
                                <div className="skeleton-img"></div>
                                <div className="skeleton-text"></div>
                            </div>
                        ))}
                </div>

                {/* {!allLoaded.current && (
                    <button className="load-more" onClick={loadMore} disabled={loading}>
                        {loading ? "Loading..." : "Load More"}
                    </button>
                )} */}
            </section>
            {selectedChar && (
                <CharacterDetails
                    character={selectedChar}
                    setSelectedChar={setSelectedChar}
                ></CharacterDetails>
            )}
        </>
    )
}

export default Home;