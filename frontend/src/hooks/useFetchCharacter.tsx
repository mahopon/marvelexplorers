import { useState, useRef } from "react";
import { Character } from "../types/character.ts";
import debounce from "../utils/debounce.ts";
import { GetCharacterResponse } from "../types/character.ts";
import { getCharacters, searchCharacter } from "../api/character.api.ts";


const useCharacterFetch = () => {
    const [characters, setCharacters] = useState<Character[]>([]);
    const [offset, setOffset] = useState(0);
    const allLoaded = useRef(false);
    const [loading, setLoading] = useState(true);
    const [filteredCharacters, setFilteredCharacters] = useState<Character[]>([]);

    const fetchCharacters = async () => {
        if (allLoaded.current) return;
        setLoading(true);
        try {
            const newChars: GetCharacterResponse[] = await getCharacters(offset)
            setTimeout(() => {
                setCharacters((prev) => [
                    ...prev,
                    ...newChars.map((char) => ({
                        id: char.ID,
                        name: char.Name,
                        description: char.Description,
                        resourceURI: char.ResourceURI,
                        thumbnailPath: char.ThumbnailPath,
                        thumbnailExtension: char.ThumbnailExtension,
                    })),
                ]);
                setLoading(false);
                setOffset(prevOffset => prevOffset + 20);
            }, 500)
        } catch {
            setLoading(false);
        }
    };

    const debouncedLoad = debounce(fetchCharacters, 500);

    const debouncedSearch = debounce(async (searchTerm: string) => {
        if (searchTerm.trim() === "") {
            setFilteredCharacters([]);
            return;
        }
        const lowercaseTerm = searchTerm.toLowerCase();
        try {
            const filteredChars: GetCharacterResponse[] = await searchCharacter(lowercaseTerm)
            setTimeout(() => {
                setFilteredCharacters((prev) => [
                    ...prev,
                    ...filteredChars.map((char) => ({
                        id: char.ID,
                        name: char.Name,
                        description: char.Description,
                        resourceURI: char.ResourceURI,
                        thumbnailPath: char.ThumbnailPath,
                        thumbnailExtension: char.ThumbnailExtension,
                    })),
                ]);
            }, 500)
        } catch (err) {
            console.error(err)
        }
    }, 500);

    const handleScroll = (containerRef: React.RefObject<HTMLDivElement | null>) => {
        if (filteredCharacters.length != 0) return false;
        const el = containerRef.current;
        if (!el || loading) return;
        const isBottom = el.scrollTop + el.clientHeight >= el.scrollHeight - 5;
        if (isBottom) debouncedLoad();
    };

    const isAtBottom = (containerRef: React.RefObject<HTMLDivElement | null>) => {
        const container = containerRef.current;
        if (!container) return false;

        const { scrollTop, scrollHeight, clientHeight } = container;
        return scrollTop + clientHeight === scrollHeight;
    };


    const checkOverflow = (containerRef: React.RefObject<HTMLDivElement | null>) => {
        console.log(filteredCharacters);
        if (filteredCharacters.length != 0) return;
        const container = containerRef.current;
        if (!container) return;

        const isOverflowing = container.scrollHeight > container.clientHeight;

        // If it's *not* overflowing OR it is at bottom of div, load more!
        if (!isOverflowing || isAtBottom(containerRef)) {
            debouncedLoad();
        }
    }

    return {
        characters,
        filteredCharacters,
        loading,
        fetchCharacters,
        debouncedSearch,
        debouncedLoad,
        checkOverflow,
        handleScroll
    };
};

export default useCharacterFetch;
