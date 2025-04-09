import { useState, useRef } from "react";
import axios from "axios";
import { Character } from "../interfaces/CharacterInterface.tsx";
import debounce from "../utils/debounce.ts";

const DATA_API = "https://tcyao.duckdns.org/api/characters?offset=";
const SEARCH_API = "https://tcyao.duckdns.org/api/characters/"

const useCharacterFetch = () => {
    const [characters, setCharacters] = useState<Character[]>([]);
    const [offset, setOffset] = useState(0);
    const allLoaded = useRef(false);
    const [loading, setLoading] = useState(true);
    const [filteredCharacters, setFilteredCharacters] = useState<Character[]>([]);
    
    const fetchCharacters = () => {
        if (allLoaded.current) return;
        setLoading(true);
        axios.get(DATA_API + offset)
            .then((res) => {
                if (res.data.length === 0) allLoaded.current = true;
                const newChars: Character[] = res.data.map((newChar: any) => ({
                    id: newChar.id,
                    name: newChar.name,
                    resourceURI: newChar.resourceuri,
                    thumbnailExtension: newChar.thumbnail_extension,
                    thumbnailPath: newChar.thumbnail_path
                }));
                setTimeout(() => {
                    setCharacters((prev) => [...prev, ...newChars]);
                    setLoading(false);
                    setOffset(prevOffset => prevOffset + 20);},500)
            }).catch((err) => {
                console.log(err);
                setLoading(false);
            });
    };

    const debouncedLoad = debounce(fetchCharacters,500);

    const debouncedSearch = debounce((searchTerm: string) => {
        if (searchTerm.trim() === "") {
            setFilteredCharacters([]);
            return;
        }
        const lowercaseTerm = searchTerm.toLowerCase();
        axios.get(SEARCH_API+lowercaseTerm)
        .then((res)=> {
            const FILTERED_CHARS: Character[] = res.data.map((char: any) => ({
                id: char.ID,
                name: char.Name,
                resourceURI: char.ResourceURI,
                thumbnailExtension: char.ThumbnailExtension,
                thumbnailPath: char.ThumbnailPath
            }));
            setFilteredCharacters(FILTERED_CHARS);
        }).catch((err)=> {
            console.log(err);
        })
    }, 500);

    const handleScroll = (containerRef: React.RefObject<HTMLDivElement | null>) => {
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
