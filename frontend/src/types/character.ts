type Character = {
    id: string;
    name: string;
    description?: string;
    resourceURI: string;
    thumbnailPath: string;
    thumbnailExtension: string;
}

type GetCharacterResponse = {
    ID: string;
    Name: string;
    Description?: string;
    ResourceURI: string;
    ThumbnailExtension: string;
    ThumbnailPath: string;
}

export type {
    Character,
    GetCharacterResponse
}
