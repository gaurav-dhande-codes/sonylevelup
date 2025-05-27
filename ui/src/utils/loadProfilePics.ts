export const loadProfilePics = (): string[] => {
  const images = import.meta.glob("../assets/profilepic/*.png", {
    eager: true,
    as: "url",
  });

  return Object.values(images);
};
