from PIL import Image, ImageDraw
import face_recognition

image = face_recognition.load_image_file("picture_library/anji.jpg")

face_landmarks_list = face_recognition.face_landmarks(image)

for face_landmarks in face_landmarks_list:
    color = [255,255,255]
    pil_image = Image.fromarray(image)
    d = ImageDraw.Draw(pil_image, 'RGBA')
    print(face_landmarks.keys())
    #           下巴，      左眉毛，       右眉毛，         鼻梁，          鼻尖,       左眼，      右眼，      上唇，      下唇
    # dict_keys(['chin', 'left_eyebrow', 'right_eyebrow', 'nose_bridge', 'nose_tip', 'left_eye', 'right_eye', 'top_lip', 'bottom_lip'])
    d.polygon(face_landmarks['left_eyebrow'], fill=(color[0], color[1], color[2], 100))
    d.polygon(face_landmarks['right_eyebrow'], fill=(color[0], color[1], color[2], 100))

    d.polygon(face_landmarks['nose_bridge'], fill=(color[0], color[1], color[2], 100))
    d.polygon(face_landmarks['nose_tip'], fill=(color[0], color[1], color[2], 100))

    d.polygon(face_landmarks['top_lip'], fill=(color[0], color[1], color[2], 100))
    d.polygon(face_landmarks['bottom_lip'], fill=(color[0], color[1], color[2], 100))

    d.line(face_landmarks['left_eye'] + [face_landmarks['left_eye'][0]], fill=(color[0], color[1], color[2], 100), width=10)
    d.line(face_landmarks['right_eye'] + [face_landmarks['right_eye'][0]], fill=(color[0], color[1], color[2], 100), width=10)

pil_image.show()
pil_image.save('picture_library/anji_cut.jpg')
