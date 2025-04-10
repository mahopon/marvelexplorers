class Character:
    def __init__(self, id, name, description, modified, thumbnail_path, thumbnail_extension, resourceuri):
        self.id = id
        self.name = name
        self.description = description
        self.modified = modified
        self.thumbnail_path = thumbnail_path
        self.thumbnail_extension = thumbnail_extension
        self.resourceuri = resourceuri
        
    def get_details(self):
        return (self.id,self.name,self.description,self.modified,self.thumbnail_path,self.thumbnail_extension,self.resourceuri)
    
    def __repr__(self):
        return (
            f"ID: {self.id}\n"
            f"Name: {self.name}\n"
            f"Description: {self.description}\n"
            f"Modified: {self.modified}\n"
            f"Thumbnail: {self.thumbnail_path}.{self.thumbnail_extension}\n"
            f"Resource URI: {self.resourceuri}"
        )