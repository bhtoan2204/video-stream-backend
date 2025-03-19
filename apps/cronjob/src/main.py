import glob
import os
import subprocess
import boto3
from collections import Counter
from PIL import Image
from transformers import pipeline

def download_video_from_s3(bucket_name, object_key, download_path):
    s3 = boto3.client('s3')
    s3.download_file(bucket_name, object_key, download_path)
    print(f"Downloaded {object_key} from bucket {bucket_name} to {download_path}")

def extract_frames_from_video(video_path, frames_dir, fps=0.2):
    if not os.path.exists(frames_dir):
        os.makedirs(frames_dir)
    command = [
        'ffmpeg',
        '-i', video_path,
        '-vf', f'fps={fps}',
        os.path.join(frames_dir, 'frame_%03d.jpg')
    ]
    print("Extracting frames from video...")
    subprocess.run(command, check=True)
    print("Frame extraction completed.")
    
def generate_tags_from_frames(frames_dir, classifier, top_k=3):
    tag_counter = Counter()
    frame_files = sorted(glob.glob(os.path.join(frames_dir, 'frame_*.jpg')))
    for frame_file in frame_files:
        print(f"Processing {frame_file}...")
        image = Image.open(frame_file).convert("RGB")
        results = classifier(image)
        for res in results[:top_k]:
            label = res['label']
            tag_counter[label] += 1

    return tag_counter

def extract_frames(video_path, frames_dir, fps=0.2):
    if not os.path.exists(frames_dir):
        os.makedirs(frames_dir)
    command = [
        'ffmpeg',
        '-i', video_path,
        '-vf', f'fps={fps}',
        os.path.join(frames_dir, 'frame_%03d.jpg')
    ]
    print("Extracting frames from video...")
    subprocess.run(command, check=True)
    print("Frame extraction completed.")

def main():
    bucket_name = 'youtube-golang' 
    object_key = 'eeabba3e-7b72-47de-a6bb-808b653ee05d/videos/20250313132042.mp4'
    video_download_path = 'video.mp4' 
    
    download_video_from_s3(bucket_name, object_key, video_download_path)
    frames_dir = 'frames'
    extract_frames(video_download_path, frames_dir, fps=0.2)
    
    classifier = pipeline("image-classification")
    
    tag_counter = generate_tags_from_frames(frames_dir, classifier, top_k=3)
    
    print("Generated tags with frequencies:")
    for tag, freq in tag_counter.most_common():
        print(f"{tag}: {freq}")
    
    final_tags = [tag for tag, freq in tag_counter.items() if freq > 1]
    print("Final tags:", final_tags)