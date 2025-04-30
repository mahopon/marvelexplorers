import api from "./api.ts"
import { GetCharacterResponse } from "../types/character.ts"
import { handleApiError } from "../utils/api-error.ts"


const getCharacters = async (offset: number) => {
    try {
        const res = await api.get<GetCharacterResponse[]>("/characters?offset=" + offset.toString())
        return res.data
    } catch (err) {
        return handleApiError(err);
    }
}

const searchCharacter = async (searchTerm: string) => {
    try {
        const res = await api.get<GetCharacterResponse[]>("/characters/" + searchTerm)
        return res.data
    } catch (err) {
        return handleApiError(err);
    }
}

export { getCharacters, searchCharacter }