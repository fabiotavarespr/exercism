import string


def is_pangram(sentence):
    # Convert sentence to lower case
    sentence = sentence.lower()

    # Create a set of ascii letter
    unused_letter = set(string.ascii_lowercase)

    # For each letter in sentence, substract that letter from unused_letter
    for letter in sentence:
        unused_letter.discard(letter)

    return len(unused_letter) == 0


