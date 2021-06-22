from PIL import Image
import face_recognition

image = face_recognition.load_image_file("picture_library/more.jpg")
face_locations = face_recognition.face_locations(image)
 
img = Image.open("picture_library/more.jpg")

k = 1
for i in face_locations:
    # 打印每张脸的位置信息
    top, right, bottom, left = i
    # print("A face is located at pixel location Top: {}, Left: {}, Bottom: {}, Right: {}".format(top, left, bottom, right))
    # 指定人脸的位置信息，然后显示人脸图片
    face_image = image[top:bottom, left:right]
    pil_image = Image.fromarray(face_image)
    # pil_image.show()
    pil_image.save("picture_library/more_cut_" + str(k) + ".jpg")
    # cropped = img.crop(i)
    # cropped.save("picture_library/Obama_eat_cut" + str(k) + ".jpg")
    k = k + 1

# pil_image.show()
# pil_image.save('picture_library/Obama_eat_cut.jpg')
