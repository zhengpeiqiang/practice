from PIL import Image

img = Image.open("picture_library/Obama_eat.jpg")
print(img.size)
cropped = img.crop((98, 85, 284, 531))  # (left, upper, right, lower) 左上，右下
# cropped = img.crop((98, 716, 284, 531))
Image._show(cropped)
cropped.save("picture_library/Obama_eat_cut.jpg")